// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// grw is the command line version of go-reader-writer (GRW) that is used for reading/writing to multiple compression/archive formats
//
package main

import (
	"fmt"
	stdos "os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	awssession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/cli"
	"github.com/spatialcurrent/go-reader-writer/pkg/grw"
	"github.com/spatialcurrent/go-reader-writer/pkg/io"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

func main() {

	rootCommand := cobra.Command{
		Use: `grw [flags] [-|stdin|INPUT_URI] [-|stdout|OUTPUT_URI]
  grw [flags] [-|stdin|INPUT_URI]
  grw [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "Read file from input and write to output",
		Long:                  "Read file from input and write to output",
		RunE: func(cmd *cobra.Command, args []string) error {
			start := time.Now()

			err := cmd.ParseFlags(args)
			if err != nil {
				return err
			}

			flag := cmd.Flags()

			v, err := cli.InitViper(flag)
			if err != nil {
				return errors.Wrap(err, "error initializing viper")
			}

			err = cli.CheckConfig(args, v)
			if err != nil {
				return err
			}

			inputUri := "stdin"
			outputUri := "stdout"
			if len(args) > 0 {
				inputUri = args[0]
				if len(args) > 1 {
					outputUri = args[1]
				}
			}

			verbose := v.GetBool(cli.FlagVerbose)

			var session *awssession.Session
			var s3Client *s3.S3

			if strings.HasPrefix(inputUri, "s3://") || strings.HasPrefix(outputUri, "s3://") {
				accessKeyID := v.GetString(cli.FlagAWSAccessKeyID)
				secretAccessKey := v.GetString(cli.FlagAWSSecretAccessKey)
				sessionToken := v.GetString(cli.FlagAWSSessionToken)

				region := v.GetString(cli.FlagAWSRegion)
				if len(region) == 0 {
					if defaultRegion := v.GetString(cli.FlagAWSDefaultRegion); len(defaultRegion) > 0 {
						region = defaultRegion
					}
				}

				config := aws.Config{
					MaxRetries: aws.Int(3),
					Region:     aws.String(region),
				}

				if len(accessKeyID) > 0 && len(secretAccessKey) > 0 {
					config.Credentials = credentials.NewStaticCredentials(
						accessKeyID,
						secretAccessKey,
						sessionToken)
				}

				session = awssession.Must(awssession.NewSessionWithOptions(awssession.Options{
					Config: config,
				}))
				s3Client = s3.New(session)
			}

			inputCompression := v.GetString(cli.FlagInputCompression)
			inputDictionary := v.GetString(cli.FlagInputDictionary)

			inputReader, _, err := grw.ReadFromResource(&grw.ReadFromResourceInput{
				Uri:        inputUri,
				Alg:        inputCompression,
				Dict:       []byte(inputDictionary),
				BufferSize: v.GetInt(cli.FlagInputBufferSize),
				S3Client:   s3Client,
			})
			if err != nil {
				return errors.Wrapf(err, "error opening resource at uri %q", inputUri)
			}

			outputCompression := v.GetString(cli.FlagOutputCompression)
			outputDictionary := v.GetString(cli.FlagOutputDictionary)
			outputOverwrite := v.GetBool(cli.FlagOutputOverwrite)
			outputAppend := v.GetBool(cli.FlagOutputAppend)
			outputBufferSize := v.GetInt(cli.FlagOutputBufferSize)

			splitLines := v.GetInt(cli.FlagSplitLines)

			var outputWriter io.ByteWriteCloser
			var outputBuffer io.Buffer

			if outputUri == "stdout" || outputUri == "-" {
				outputWriter, err = grw.WriteStdout(outputCompression, []byte(outputDictionary))
				if err != nil {
					return errors.Wrap(err, "error opening stdout")
				}
			} else if strings.HasPrefix(outputUri, "s3://") {
				outputWriter, outputBuffer, err = grw.WriteBytes(outputCompression, []byte(outputDictionary))
				if err != nil {
					return errors.Wrapf(err, "error opening bytes buffer for %q", outputUri)
				}
			} else {
				uri := outputUri
				if splitLines > 0 {
					uri = strings.ReplaceAll(outputUri, cli.NumberReplacementCharacter, "1")
				}
				if (!outputOverwrite) && (!outputAppend) {
					exists, _, err := os.Stat(uri)
					if err != nil {
						return errors.Wrapf(err, "error statting uri %q", uri)
					}
					if exists {
						return fmt.Errorf("file already exists at uri %q and neither append or overwrite is set", uri)
					}
				}
				outputWriter, err = grw.WriteToResource(&grw.WriteToResourceInput{
					Uri:      uri,
					Alg:      outputCompression,
					Dict:     []byte(outputDictionary),
					Append:   outputAppend,
					S3Client: s3Client,
				})
				if err != nil {
					return errors.Wrapf(err, "error opening resource at uri %q", outputUri)
				}
			}

			var wg sync.WaitGroup
			wg.Add(1)
			signals := make(chan stdos.Signal, 1)
			signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGPIPE)

			gracefulMutex := &sync.Mutex{}
			gracefulShutdown := false

			updateGracefulShutdown := func(value *bool) bool {
				gracefulMutex.Lock()
				if value != nil {
					gracefulShutdown = *value
				} else {
					value = &gracefulShutdown
				}
				gracefulMutex.Unlock()
				return *value
			}

			go func() {
				<-signals
				value := true
				updateGracefulShutdown(&value)
			}()

			brokenPipe := false
			if splitLines > 0 {
				go func() {
					eof := false

					scanner := bufio.NewScanner(inputReader)
					files := 1
					lines := 0

					for (!updateGracefulShutdown(nil)) && (!eof) && (!brokenPipe) && scanner.Scan() {

						if lines == splitLines {

							err := outputWriter.Flush()
							if err != nil {
								fmt.Fprint(os.Stderr, errors.Wrapf(err, "error flushing resource at uri %q", strings.ReplaceAll(outputUri, cli.NumberReplacementCharacter, strconv.Itoa(files))).Error())
								break
							}

							err = outputWriter.Close()
							if err != nil {
								fmt.Fprint(os.Stderr, errors.Wrapf(err, "error closing resource at uri %q", strings.ReplaceAll(outputUri, cli.NumberReplacementCharacter, strconv.Itoa(files))).Error())
								break
							}

							// increment files number
							files++

							uri := strings.ReplaceAll(outputUri, cli.NumberReplacementCharacter, strconv.Itoa(files))

							if (!outputOverwrite) && (!outputAppend) {
								exists, _, err := os.Stat(uri)
								if err != nil {
									fmt.Fprint(os.Stderr, errors.Wrapf(err, "error statting uri %q", uri).Error())
									break
								}
								if exists {
									fmt.Fprint(os.Stderr, fmt.Errorf("file already exists at uri %q and neither append or overwrite is set", uri).Error())
									break
								}
							}

							ow, err := grw.WriteToResource(&grw.WriteToResourceInput{
								Uri:      uri,
								Alg:      outputCompression,
								Dict:     []byte(outputDictionary),
								Append:   outputAppend,
								S3Client: s3Client,
							})
							if err != nil {
								fmt.Fprint(os.Stderr, errors.Wrapf(err, "error opening resource at uri %q", outputUri).Error())
								break
							}

							outputWriter = ow

							lines = 0
						}

						line := scanner.Text()

						if gracefulShutdown {
							break
						}

						_, err = outputWriter.WriteLine(line)
						if err != nil {
							if perr, ok := err.(*stdos.PathError); ok {
								if perr.Err == syscall.EPIPE {
									brokenPipe = true
									break
								}
							}
							fmt.Fprint(os.Stderr, errors.Wrapf(err, "error writing to resource at uri %q", outputUri).Error())
							break
						}

						// increment counter
						lines++
					}

					if err := scanner.Err(); err != nil {
						fmt.Fprint(os.Stderr, errors.Wrapf(err, "error reading from resource at uri %q", inputUri).Error())
					}

					wg.Done()
				}()
			} else {
				go func() {
					eof := false
					for (!updateGracefulShutdown(nil)) && (!eof) && (!brokenPipe) {

						b := make([]byte, outputBufferSize)
						n, err := inputReader.Read(b)
						if err != nil {
							if err == io.EOF {
								eof = true
								// do not break
								// if the input is less than the size of the buffer,
								// will then use n > 0, n < len(b), and return EOF
							} else {
								fmt.Fprintln(os.Stderr, errors.Wrapf(err, "error reading from resource at uri %q", inputUri).Error())
								break
							}
						}

						if gracefulShutdown {
							break
						}

						if n > 0 {
							_, err = outputWriter.Write(b[:n])
							if err != nil {
								if perr, ok := err.(*stdos.PathError); ok {
									if perr.Err == syscall.EPIPE {
										brokenPipe = true
										break
									}
								}
								fmt.Fprintln(os.Stderr, errors.Wrapf(err, "error writing to resource at uri %q", outputUri).Error())
							}
						}

					}
					wg.Done()
				}()
			}

			wg.Wait() // wait until done writing or received signal for graceful shutdown

			errorReader, errorWriter := grw.CloseReaderAndWriter(inputReader, outputWriter, brokenPipe)
			if errorReader != nil || errorWriter != nil {
				if errorReader != nil {
					fmt.Fprint(os.Stderr, errorReader.Error())
				}
				if errorWriter != nil {
					fmt.Fprint(os.Stderr, errorWriter.Error())
				}
				os.Exit(1)
			}

			if strings.HasPrefix(outputUri, "s3://") {
				_, outputPath := splitter.SplitUri(outputUri)
				i := strings.Index(outputPath, "/")
				if i == -1 {
					return errors.Wrap(err, "path missing bucket")
				}
				err := grw.UploadS3Object(outputPath[0:i], outputPath[i+1:], outputBuffer, s3Client)
				if err != nil {
					return errors.Wrapf(err, "error uploading to AWS S3 at uri %q", outputUri)
				}
			}

			elapsed := time.Since(start)
			if verbose && !brokenPipe {
				fmt.Println("Done in " + elapsed.String())
			}

			return nil
		},
	}
	cli.InitFlags(rootCommand.Flags())

	if err := rootCommand.Execute(); err != nil {
		panic(err)
	}
}

// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// grw is the command line version of go-reader-writer (GRW) that is used for reading/writing to multiple compression/archive formats
//
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	stdos "os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/spf13/cobra"

	awssession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/cli"
	"github.com/spatialcurrent/go-reader-writer/pkg/grw"
	"github.com/spatialcurrent/go-reader-writer/pkg/io"
	"github.com/spatialcurrent/go-reader-writer/pkg/nop"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

const (
	GRWVersion = "v0.0.3"
)

func initPrivateKey(p string) ([]byte, error) {
	if len(p) == 0 {
		return []byte{}, nil
	}
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("error reading private key from %q: %w", p, err)
	}
	return b, nil
}

func main() {

	rootCommand := cobra.Command{
		Use: `grw [flags] [-|INPUT_URI] [-|OUTPUT_URI]
  grw [flags] [-|INPUT_URI]
  grw [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "grw is a simple tool for reading and writing compressed resources by uri.",
		Long: `grw is a simple tool for reading and writing compressed resources by uri.
By default, reads from stdin and writes to stdout.
If the output uri is a device, then the append flag is not required.
Supports the following compression algorithms: ` + strings.Join(grw.Algorithms, ", "),
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			start := time.Now()

			err := cmd.ParseFlags(args)
			if err != nil {
				return err
			}

			flag := cmd.Flags()

			v, err := cli.InitViper(flag)
			if err != nil {
				return fmt.Errorf("error initializing viper: %w", err)
			}

			if v.GetBool(cli.FlagVersion) {
				fmt.Println(GRWVersion)
				return nil
			}

			err = cli.CheckConfig(args, v)
			if err != nil {
				return err
			}

			inputUri := "-"
			outputUri := "-"
			if len(args) > 0 {
				inputUri = args[0]
				if len(args) > 1 {
					outputUri = args[1]
				}
			}

			verbose := v.GetBool(cli.FlagVerbose)

			outputBufferSize := v.GetInt(cli.FlagOutputBufferSize)
			if outputBufferSize < 0 {
				if outputUri == "-" {
					outputBufferSize = 4096 // the buffer size is used when reading from stdin and not for the writer.
				} else {
					outputBufferSize = 4096
				}
			}

			inputPrivateKey, err := initPrivateKey(v.GetString(cli.FlagInputPrivateKey))
			if err != nil {
				return fmt.Errorf("error initializing input private key: %w", err)
			}

			outputPrivateKey, err := initPrivateKey(v.GetString(cli.FlagOutputPrivateKey))
			if err != nil {
				return fmt.Errorf("error initializing output private key: %w", err)
			}

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

			if strings.HasPrefix(inputUri, "s3://") || strings.HasPrefix(outputUri, "s3://") {

			}

			inputCompression := v.GetString(cli.FlagInputCompression)
			inputDictionary := v.GetString(cli.FlagInputDictionary)

			if inputUri != "-" {
				exists, fileInfo, err := os.Stat(inputUri)
				if err != nil {
					return fmt.Errorf("error stating resource at uri %q: %w", inputUri, err)
				}

				if !exists {
					return fmt.Errorf("resource at input uri %q does not exist", inputUri)
				}

				if !(fileInfo.IsRegular() || fileInfo.IsNamedPipe()) {
					return fmt.Errorf("resource at input uri %q is neither a regular file or named pipe: %w", inputUri, err)
				}
			}

			readFromResourceOutput, err := grw.ReadFromResource(&grw.ReadFromResourceInput{
				URI:        inputUri,
				Alg:        inputCompression,
				Dict:       []byte(inputDictionary),
				BufferSize: v.GetInt(cli.FlagInputBufferSize),
				S3Client:   s3Client,
				PrivateKey: inputPrivateKey,
			})
			if err != nil {
				return fmt.Errorf("error opening resource at uri %q: %w", inputUri, err)
			}
			inputReader := readFromResourceOutput.Reader

			outputCompression := v.GetString(cli.FlagOutputCompression)
			outputDictionary := v.GetString(cli.FlagOutputDictionary)
			outputOverwrite := v.GetBool(cli.FlagOutputOverwrite)
			outputAppend := v.GetBool(cli.FlagOutputAppend)

			splitLines := v.GetInt(cli.FlagSplitLines)

			var outputWriter io.WriteCloser
			var outputBuffer io.Buffer

			if outputUri == "-" {
				outputWriter, err = grw.WrapWriter(nop.NewWriteCloser(os.Stdout), outputCompression, []byte(outputDictionary), grw.NoBuffer)
				if err != nil {
					return fmt.Errorf("error opening stdout: %w", err)
				}
			} else if strings.HasPrefix(outputUri, "s3://") {
				outputBuffer = new(bytes.Buffer)
				outputWriter, err = grw.WrapWriter(nop.NewWriteCloser(outputBuffer), outputCompression, []byte(outputDictionary), grw.NoBuffer)
				if err != nil {
					return fmt.Errorf("error opening bytes buffer for %q: %w", outputUri, err)
				}
			} else {
				uri := outputUri
				if splitLines > 0 {
					uri = strings.ReplaceAll(outputUri, cli.NumberReplacementCharacter, "1")
				}
				scheme, path := splitter.SplitUri(uri)
				if scheme == "sftp" {
					writeToResourceOutput, err := grw.WriteToResource(&grw.WriteToResourceInput{
						URI:        uri,
						Alg:        outputCompression,
						BufferSize: outputBufferSize,
						Dict:       []byte(outputDictionary),
						Append:     outputAppend,
						S3Client:   s3Client,
						PrivateKey: outputPrivateKey,
					})
					if err != nil {
						return fmt.Errorf("error writing to resource at uri %q: %w", outputUri, err)
					}
					outputWriter = writeToResourceOutput.Writer
				} else if scheme == "file" || scheme == "" {
					if (!outputOverwrite) && (!outputAppend) {
						exists, fileInfo, err := os.Stat(path)
						if err != nil {
							return fmt.Errorf("error statting uri %q: %w", uri, err)
						}
						if exists && (!fileInfo.IsDevice()) && (!fileInfo.IsNamedPipe()) {
							return fmt.Errorf("file already exists at uri %q and neither append or overwrite is set", uri)
						}
					}
					if v.GetBool(cli.FlagOutputMkdirs) {
						exists, _, err := os.Stat(filepath.Dir(path))
						if err != nil {
							return fmt.Errorf("error statting uri %q: %w", uri, err)
						}
						if !exists {
							err := os.MkdirAll(filepath.Dir(path), 0770)
							if err != nil {
								return fmt.Errorf("error creating parent directories for uri %q: %w", uri, err)
							}
						}
					}
					writeToResourceOutput, err := grw.WriteToResource(&grw.WriteToResourceInput{
						URI:        uri,
						Alg:        outputCompression,
						BufferSize: outputBufferSize,
						Dict:       []byte(outputDictionary),
						Append:     outputAppend,
						S3Client:   s3Client,
						PrivateKey: outputPrivateKey,
					})
					if err != nil {
						return fmt.Errorf("error writing to resource at uri %q: %w", outputUri, err)
					}
					outputWriter = writeToResourceOutput.Writer
				} else {
					return errors.New(fmt.Sprintf("unknown scheme for uri %q", outputUri))
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

							if outputFlusher, ok := outputWriter.(interface{ Flush() error }); ok {
								err := outputFlusher.Flush()
								if err != nil {
									fmt.Fprint(os.Stderr, fmt.Errorf("error flushing resource at uri %q: %w", strings.ReplaceAll(outputUri, cli.NumberReplacementCharacter, strconv.Itoa(files)), err).Error())
									break
								}
							}

							err = outputWriter.Close()
							if err != nil {
								fmt.Fprint(os.Stderr, fmt.Errorf("error closing resource at uri %q: %w", strings.ReplaceAll(outputUri, cli.NumberReplacementCharacter, strconv.Itoa(files)), err).Error())
								break
							}

							// increment files number
							files++

							uri := strings.ReplaceAll(outputUri, cli.NumberReplacementCharacter, strconv.Itoa(files))

							scheme, path := splitter.SplitUri(uri)
							if scheme == "file" || scheme == "" {
								if (!outputOverwrite) && (!outputAppend) {
									exists, fileInfo, err := os.Stat(path)
									if err != nil {
										fmt.Fprint(os.Stderr, fmt.Errorf("error statting uri %q: %w", uri, err).Error())
										break
									}
									if exists && (!fileInfo.IsDevice()) && (!fileInfo.IsNamedPipe()) {
										fmt.Fprintln(os.Stderr, fmt.Errorf("file already exists at uri %q and neither append or overwrite is set", uri).Error())
										break
									}
								}
								if v.GetBool(cli.FlagOutputMkdirs) {
									exists, _, err := os.Stat(filepath.Dir(path))
									if err != nil {
										fmt.Fprint(os.Stderr, fmt.Errorf("error statting uri %q: %w", uri, err).Error())
										break
									}
									if !exists {
										err := os.MkdirAll(filepath.Dir(path), 0770)
										if err != nil {
											fmt.Fprint(os.Stderr, fmt.Errorf("error creating parent directories for uri %q: %w", uri, err).Error())
											break
										}
									}
								}
								writeToResourceOutput, err := grw.WriteToResource(&grw.WriteToResourceInput{
									URI:        uri,
									Alg:        outputCompression,
									BufferSize: outputBufferSize,
									Dict:       []byte(outputDictionary),
									Append:     outputAppend,
									S3Client:   s3Client,
									PrivateKey: outputPrivateKey,
								})
								if err != nil {
									fmt.Fprint(os.Stderr, fmt.Errorf("error opening resource at uri %q: %w", outputUri, err).Error())
									break
								}
								outputWriter = writeToResourceOutput.Writer
							} else {
								fmt.Fprintf(os.Stderr, "unknown scheme for uri %q", outputUri)
								break
							}

							lines = 0
						}

						line := scanner.Text()

						if gracefulShutdown {
							break
						}

						_, err = io.WriteLine(outputWriter, line)
						if err != nil {
							if perr, ok := err.(*stdos.PathError); ok {
								if perr.Err == syscall.EPIPE {
									brokenPipe = true
									break
								}
							}
							fmt.Fprint(os.Stderr, fmt.Errorf("error writing to resource at uri %q: %w", outputUri, err).Error())
							break
						}

						// increment counter
						lines++
					}

					if err := scanner.Err(); err != nil {
						fmt.Fprint(os.Stderr, fmt.Errorf("error reading from resource at uri %q: %w", inputUri, err).Error())
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
								fmt.Fprintln(os.Stderr, fmt.Errorf("error reading from resource at uri %q: %w", inputUri, err).Error())
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
								fmt.Fprintln(os.Stderr, fmt.Errorf("error writing to resource at uri %q: %w", outputUri, err).Error())
							}
						}

					}
					wg.Done()
				}()
			}

			wg.Wait() // wait until done writing or received signal for graceful shutdown

			errorReader, errorWriter := io.CloseReaderAndWriter(inputReader, outputWriter, brokenPipe)
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
					return fmt.Errorf("path missing bucket: %w", err)
				}
				err := grw.UploadS3Object(outputPath[0:i], outputPath[i+1:], outputBuffer, s3Client)
				if err != nil {
					return fmt.Errorf("error uploading to AWS S3 at uri %q: %w", outputUri, err)
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
		fmt.Fprintln(os.Stderr, "grw: "+err.Error())
		fmt.Fprintln(os.Stderr, "Try grw --help for more information.")
		os.Exit(1)
	}
}

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
	"io"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	awssession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/spatialcurrent/go-reader-writer/pkg/grw"
	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

const (
	flagAWSProfile         string = "aws-profile"
	flagAWSDefaultRegion   string = "aws-default-region"
	flagAWSRegion          string = "aws-region"
	flagAWSAccessKeyID     string = "aws-access-key-id"
	flagAWSSecretAccessKey string = "aws-secret-access-key"
	flagAWSSessionToken    string = "aws-session-token"
	flagInputCompression   string = "input-compression"
	flagInputBufferSize    string = "input-buffer-size"
	flagOutputCompression  string = "output-compression"
	flagOutputBufferSize   string = "output-buffer-size"
	flagOutputAppend       string = "output-append"
	flagOutputOverwrite    string = "output-overwrite"
	flagVerbose            string = "verbose"
)

func initFlags(flag *pflag.FlagSet) {
	flag.String(flagAWSProfile, "", "AWS Profile")
	flag.String(flagAWSDefaultRegion, "", "AWS Default Region")
	flag.StringP(flagAWSRegion, "", "", "AWS Region (overrides default region)")
	flag.StringP(flagAWSAccessKeyID, "", "", "AWS Access Key ID")
	flag.StringP(flagAWSSecretAccessKey, "", "", "AWS Secret Access Key")
	flag.StringP(flagAWSSessionToken, "", "", "AWS Session Token")

	flag.StringP(flagInputCompression, "", "", "the input compression: "+strings.Join(grw.Algorithms, ", "))
	flag.Int(flagInputBufferSize, 4096, "the input reader buffer size")

	flag.StringP(flagOutputCompression, "", "", "the output compression: "+strings.Join(grw.Algorithms, ", "))
	flag.IntP(flagOutputBufferSize, "b", 4096, "the output writer buffer size")
	flag.BoolP(flagOutputAppend, "a", false, "append to output files")
	flag.BoolP(flagOutputOverwrite, "o", false, "overwrite output if it already exists")

	flag.BoolP(flagVerbose, "v", false, "verbose output")
}

func main() {

	rootCommand := cobra.Command{
		Use:                   "grw [flags] [-|stdin|INPUT_URI] [-|stdout|OUTPUT_URI]",
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

			v := viper.New()
			err = v.BindPFlags(flag)
			if err != nil {
				return err
			}
			v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
			v.AutomaticEnv()

			if len(args) != 2 {
				return fmt.Errorf("expecting 2 arguments, found %d arguments", len(args))
			}

			inputUri := args[0]
			outputUri := args[1]

			verbose := v.GetBool(flagVerbose)

			var session *awssession.Session
			var s3Client *s3.S3

			if strings.HasPrefix(inputUri, "s3://") || strings.HasPrefix(outputUri, "s3://") {
				accessKeyID := v.GetString(flagAWSAccessKeyID)
				secretAccessKey := v.GetString(flagAWSSecretAccessKey)
				sessionToken := v.GetString(flagAWSSessionToken)

				region := v.GetString(flagAWSRegion)
				if len(region) == 0 {
					if defaultRegion := v.GetString(flagAWSDefaultRegion); len(defaultRegion) > 0 {
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

			inputCompression := v.GetString(flagInputCompression)

			inputReader, _, err := grw.ReadFromResource(
				inputUri,
				inputCompression,
				v.GetInt(flagInputBufferSize),
				s3Client,
			)
			if err != nil {
				return errors.Wrapf(err, "error opening resource at uri %q", inputUri)
			}

			outputCompression := v.GetString(flagOutputCompression)
			outputAppend := v.GetBool(flagOutputAppend)
			outputBufferSize := v.GetInt(flagOutputBufferSize)

			var outputWriter grw.ByteWriteCloser
			var outputBuffer grw.Buffer

			if outputUri == "stdout" || outputUri == "-" {
				outputWriter, err = grw.WriteStdout(outputCompression)
				if err != nil {
					return errors.Wrap(err, "error opening stdout")
				}
			} else if strings.HasPrefix(outputUri, "s3://") {
				outputWriter, outputBuffer, err = grw.WriteBytes(outputCompression)
				if err != nil {
					return errors.Wrapf(err, "error opening bytes buffer for %q", outputUri)
				}
			} else {
				outputWriter, err = grw.WriteToResource(outputUri, outputCompression, outputAppend, s3Client)
				if err != nil {
					return errors.Wrapf(err, "error opening resource at uri %q", outputUri)
				}
			}

			var wg sync.WaitGroup
			wg.Add(1)
			signals := make(chan os.Signal, 1)
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
			go func() {
				eof := false
				for (!updateGracefulShutdown(nil)) && (!eof) && (!brokenPipe) {

					b := make([]byte, outputBufferSize)
					n, err := inputReader.Read(b)
					if err != nil {
						if err == io.EOF {
							eof = true
						} else {
							fmt.Fprint(os.Stderr, errors.Wrapf(err, "error reading from resource at uri %q", inputUri).Error())
						}
					}

					if gracefulShutdown {
						break
					}

					_, err = outputWriter.Write(b[:n])
					if err != nil {
						if perr, ok := err.(*os.PathError); ok {
							if perr.Err == syscall.EPIPE {
								brokenPipe = true
								break
							}
						}
						fmt.Fprint(os.Stderr, errors.Wrapf(err, "error writing to resource at uri %q", outputUri).Error())
					}

				}
				wg.Done()
			}()

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
	initFlags(rootCommand.Flags())

	if err := rootCommand.Execute(); err != nil {
		panic(err)
	}
}

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
	"fmt"
	stdos "os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/pkg/sftp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	awssession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/cli"
	"github.com/spatialcurrent/go-reader-writer/pkg/grw"
	"github.com/spatialcurrent/go-reader-writer/pkg/io"
	"github.com/spatialcurrent/go-reader-writer/pkg/net/sftp2"
	"github.com/spatialcurrent/go-reader-writer/pkg/net/ssh2"
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
	b, err := stdos.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("error reading private key from %q: %w", p, err)
	}
	return b, nil
}

func initPrivateKeys(inputPath string, outputPath string) ([]byte, []byte, error) {
	inputPrivateKey, err := initPrivateKey(inputPath)
	if err != nil {
		return nil, nil, fmt.Errorf("error initializing input private key: %w", err)
	}

	outputPrivateKey, err := initPrivateKey(outputPath)
	if err != nil {
		return inputPrivateKey, nil, fmt.Errorf("error initializing output private key: %w", err)
	}

	return inputPrivateKey, outputPrivateKey, nil
}

func initS3Client(v *viper.Viper, inputURI string, outputURI string) (*s3.S3, *awssession.Session, error) {
	if (!strings.HasPrefix(inputURI, "s3://")) && (!strings.HasPrefix(outputURI, "s3://")) {
		return nil, nil, nil
	}

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

	session, err := awssession.NewSessionWithOptions(awssession.Options{
		Config: config,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("error creating new AWS session: %w", err)
	}
	return s3.New(session), session, nil
}

func initSFTPClient(uri string, password string, privateKeyBytes []byte) (*sftp.Client, *ssh.Client, error) {
	if !strings.HasPrefix(uri, "sftp://") {
		return nil, nil, nil
	}

	options := []ssh2.ClientOption{}

	if len(privateKeyBytes) > 0 {
		privateKey, err := ssh.ParsePrivateKey(privateKeyBytes)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing private key: %w", err)
		}
		options = append(options, func(config *ssh2.ClientConfig) error {
			config.Auth = []ssh.AuthMethod{
				ssh.PublicKeys(privateKey),
			}
			return nil
		})
	} else if len(password) > 0 {
		options = append(options, func(config *ssh2.ClientConfig) error {
			config.Auth = []ssh.AuthMethod{
				ssh.Password(password),
			}
			return nil
		})
	}

	sshClient, err := ssh2.Dial(uri, options...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating SSH client: %w", err)
	}

	sftpClient, err := sftp.NewClient(sshClient.Client)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating SFTP client: %w", err)
	}

	return sftpClient, sshClient.Client, nil
}

func checkURIRead(uri string, sftpClient *sftp.Client) error {
	if uri != "-" {
		scheme, path := splitter.SplitURI(uri)
		switch scheme {
		case "sftp":
			err := sftp2.CheckFileRead(sftpClient, strings.SplitN(path, "/", 2)[1])
			if err != nil {
				return fmt.Errorf("resource %q cannot be read: %w", uri, err)
			}
		case "file", "":
			err := os.CheckURIRead(uri)
			if err != nil {
				return fmt.Errorf("resource %q cannot be read: %w", uri, err)
			}
		}
	}
	return nil
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

			inputURI := "-"
			outputURI := "-"
			if len(args) > 0 {
				inputURI = args[0]
				if len(args) > 1 {
					outputURI = args[1]
				}
			}

			_, outputPath := splitter.SplitURI(outputURI)

			verbose := v.GetBool(cli.FlagVerbose)

			outputModeString := v.GetString(cli.FlagOutputMode)

			outputMode, outputModeErr := strconv.ParseUint(outputModeString, 0, 32)
			if outputModeErr != nil {
				return fmt.Errorf("invalid output mode %q: %w", outputModeString, outputModeErr)
			}

			outputACL := v.GetString(cli.FlagOutputACL)

			fmt.Println("output acl:", outputACL)

			outputBufferSize := v.GetInt(cli.FlagOutputBufferSize)
			if outputBufferSize < 0 {
				if outputURI == "-" {
					outputBufferSize = 4096 // the buffer size is used when reading from stdin and not for the writer.
				} else {
					outputBufferSize = 4096
				}
			}

			inputPassword := v.GetString(cli.FlagInputPassword)

			outputPassword := v.GetString(cli.FlagOutputPassword)

			inputPrivateKey, outputPrivateKey, err := initPrivateKeys(v.GetString(cli.FlagInputPrivateKey), v.GetString(cli.FlagOutputPrivateKey))
			if err != nil {
				return fmt.Errorf("error initializing private keys: %w", err)
			}

			s3Client, _, err := initS3Client(v, inputURI, outputURI)
			if err != nil {
				return fmt.Errorf("error initializing AWS S3 client: %w", err)
			}

			inputSFTPClient, inputSSHClient, err := initSFTPClient(inputURI, inputPassword, inputPrivateKey)
			if err != nil {
				return fmt.Errorf("error initializing SFTP client for input at %q: %w", inputURI, err)
			}

			outputSFTPClient, outputSSHClient, err := initSFTPClient(outputURI, outputPassword, outputPrivateKey)
			if err != nil {
				return fmt.Errorf("error initializing SFTP client for output at %q: %w", outputURI, err)
			}

			inputCompression := v.GetString(cli.FlagInputCompression)
			inputDictionary := v.GetString(cli.FlagInputDictionary)

			err = checkURIRead(inputURI, inputSFTPClient)
			if err != nil {
				_ = inputSFTPClient.Close()
				_ = inputSSHClient.Close()
				return fmt.Errorf("input resource %q not valid: %w", inputURI, err)
			}

			readFromResourceOutput, err := grw.ReadFromResource(&grw.ReadFromResourceInput{
				URI:        inputURI,
				Alg:        inputCompression,
				Dict:       []byte(inputDictionary),
				BufferSize: v.GetInt(cli.FlagInputBufferSize),
				S3Client:   s3Client,
				SSHClient:  inputSSHClient,
				SFTPClient: inputSFTPClient,
				Password:   inputPassword,
				PrivateKey: inputPrivateKey,
			})
			if err != nil {
				return fmt.Errorf("error opening resource at uri %q: %w", inputURI, err)
			}
			inputReader := readFromResourceOutput.Reader

			outputCompression := v.GetString(cli.FlagOutputCompression)
			outputDictionary := v.GetString(cli.FlagOutputDictionary)
			outputOverwrite := v.GetBool(cli.FlagOutputOverwrite)
			outputAppend := v.GetBool(cli.FlagOutputAppend)

			splitLines := v.GetInt(cli.FlagSplitLines)

			var outputWriter io.WriteCloser
			var outputBuffer io.Buffer

			if outputURI == "-" {
				outputWriter, err = grw.WrapWriter(nop.NewWriteCloser(os.Stdout), outputCompression, []byte(outputDictionary), grw.NoBuffer)
				if err != nil {
					return fmt.Errorf("error opening stdout: %w", err)
				}
			} else if strings.HasPrefix(outputURI, "s3://") {
				outputBuffer = new(bytes.Buffer)
				outputWriter, err = grw.WrapWriter(nop.NewWriteCloser(outputBuffer), outputCompression, []byte(outputDictionary), grw.NoBuffer)
				if err != nil {
					return fmt.Errorf("error opening bytes buffer for %q: %w", outputURI, err)
				}
			} else {
				uri := outputURI
				if splitLines > 0 {
					uri = strings.ReplaceAll(outputURI, cli.NumberReplacementCharacter, "1")
				}
				scheme, path := splitter.SplitURI(uri)
				if scheme == "sftp" {
					err = sftp2.CheckFileWrite(outputSFTPClient, strings.SplitN(path, "/", 2)[1], outputAppend, outputOverwrite)
					if err != nil {
						return fmt.Errorf("cannot write to resource at uri %q: %w", outputURI, err)
					}
					writeToResourceOutput, errWriteToResource := grw.WriteToResource(&grw.WriteToResourceInput{
						ACL:        outputACL,
						Append:     outputAppend,
						Alg:        outputCompression,
						BufferSize: outputBufferSize,
						Dict:       []byte(outputDictionary),
						Mode:       uint32(outputMode),
						Password:   outputPassword,
						PrivateKey: outputPrivateKey,
						S3Client:   s3Client,
						SSHClient:  outputSSHClient,
						SFTPClient: outputSFTPClient,
						URI:        uri,
					})
					if errWriteToResource != nil {
						return fmt.Errorf("error writing to resource at uri %q: %w", outputURI, errWriteToResource)
					}
					outputWriter = writeToResourceOutput.Writer
				} else if scheme == "file" || scheme == "" {
					err = os.CheckURIWrite(uri, outputAppend, outputOverwrite)
					if err != nil {
						return fmt.Errorf("cannot write to resource at uri %q: %w", outputURI, err)
					}
					if v.GetBool(cli.FlagOutputMkdirs) {
						exists, _, errStat := os.Stat(filepath.Dir(path))
						if err != nil {
							return fmt.Errorf("error statting uri %q: %w", uri, errStat)
						}
						if !exists {
							err = os.MkdirAll(filepath.Dir(path), 0770)
							if err != nil {
								return fmt.Errorf("error creating parent directories for uri %q: %w", uri, err)
							}
						}
					}
					writeToResourceOutput, errWriteToResource := grw.WriteToResource(&grw.WriteToResourceInput{
						ACL:        outputACL,
						Alg:        outputCompression,
						Append:     outputAppend,
						BufferSize: outputBufferSize,
						Dict:       []byte(outputDictionary),
						Mode:       uint32(outputMode),
						Password:   outputPassword,
						PrivateKey: outputPrivateKey,
						S3Client:   s3Client,
						SSHClient:  outputSSHClient,
						SFTPClient: outputSFTPClient,
						URI:        uri,
					})
					if errWriteToResource != nil {
						return fmt.Errorf("error writing to resource at uri %q: %w", outputURI, errWriteToResource)
					}
					outputWriter = writeToResourceOutput.Writer
				} else {
					return fmt.Errorf("unknown scheme for uri %q", outputURI)
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
								errFlush := outputFlusher.Flush()
								if errFlush != nil {
									fmt.Fprint(os.Stderr, fmt.Errorf("error flushing resource at uri %q: %w", strings.ReplaceAll(outputURI, cli.NumberReplacementCharacter, strconv.Itoa(files)), errFlush).Error())
									break
								}
							}

							errClose := outputWriter.Close()
							if errClose != nil {
								fmt.Fprint(os.Stderr, fmt.Errorf("error closing resource at uri %q: %w", strings.ReplaceAll(outputURI, cli.NumberReplacementCharacter, strconv.Itoa(files)), errClose).Error())
								break
							}

							// increment files number
							files++

							uri := strings.ReplaceAll(outputURI, cli.NumberReplacementCharacter, strconv.Itoa(files))

							scheme, path := splitter.SplitURI(uri)
							if scheme == "sftp" {
								errCheckFileWrite := sftp2.CheckFileWrite(outputSFTPClient, strings.SplitN(path, "/", 2)[1], outputAppend, outputOverwrite)
								if errCheckFileWrite != nil {
									fmt.Fprint(os.Stderr, fmt.Errorf("cannot write to resource at uri %q: %w", uri, errCheckFileWrite).Error())
									break
								}
								writeToResourceOutput, errWriteToResource := grw.WriteToResource(&grw.WriteToResourceInput{
									ACL:        outputACL,
									Alg:        outputCompression,
									Append:     outputAppend,
									BufferSize: outputBufferSize,
									Dict:       []byte(outputDictionary),
									Mode:       uint32(outputMode),
									Password:   outputPassword,
									PrivateKey: outputPrivateKey,
									S3Client:   s3Client,
									SSHClient:  outputSSHClient,
									SFTPClient: outputSFTPClient,
									URI:        uri,
								})
								if errWriteToResource != nil {
									fmt.Fprint(os.Stderr, fmt.Errorf("error opening resource at uri %q: %w", outputURI, errWriteToResource).Error())
									break
								}
								outputWriter = writeToResourceOutput.Writer
							} else if scheme == "file" || scheme == "" {
								errCheckURIWrite := os.CheckURIWrite(uri, outputAppend, outputOverwrite)
								if errCheckURIWrite != nil {
									fmt.Fprint(os.Stderr, fmt.Errorf("cannot write to resource at uri %q: %w", uri, errCheckURIWrite).Error())
									break
								}
								if v.GetBool(cli.FlagOutputMkdirs) {
									exists, _, errStat := os.Stat(filepath.Dir(path))
									if errStat != nil {
										fmt.Fprint(os.Stderr, fmt.Errorf("error statting uri %q: %w", uri, errStat).Error())
										break
									}
									if !exists {
										errMkdirAll := os.MkdirAll(filepath.Dir(path), 0770)
										if errMkdirAll != nil {
											fmt.Fprint(os.Stderr, fmt.Errorf("error creating parent directories for uri %q: %w", uri, errMkdirAll).Error())
											break
										}
									}
								}
								writeToResourceOutput, errWriteToResource := grw.WriteToResource(&grw.WriteToResourceInput{
									ACL:        outputACL,
									Alg:        outputCompression,
									Append:     outputAppend,
									BufferSize: outputBufferSize,
									Dict:       []byte(outputDictionary),
									Mode:       uint32(outputMode),
									Password:   outputPassword,
									PrivateKey: outputPrivateKey,
									S3Client:   s3Client,
									SSHClient:  outputSSHClient,
									SFTPClient: outputSFTPClient,
									URI:        uri,
								})
								if errWriteToResource != nil {
									fmt.Fprint(os.Stderr, fmt.Errorf("error opening resource at uri %q: %w", outputURI, errWriteToResource).Error())
									break
								}
								outputWriter = writeToResourceOutput.Writer
							} else {
								fmt.Fprintf(os.Stderr, "unknown scheme for uri %q", outputURI)
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
							fmt.Fprint(os.Stderr, fmt.Errorf("error writing to resource at uri %q: %w", outputURI, err).Error())
							break
						}

						// increment counter
						lines++
					}

					if errScan := scanner.Err(); errScan != nil {
						fmt.Fprint(os.Stderr, fmt.Errorf("error reading from resource at uri %q: %w", inputURI, errScan).Error())
					}

					wg.Done()
				}()
			} else {
				go func() {
					eof := false
					for (!updateGracefulShutdown(nil)) && (!eof) && (!brokenPipe) {

						b := make([]byte, outputBufferSize)
						n, errRead := inputReader.Read(b)
						if errRead != nil {
							if errRead == io.EOF {
								eof = true
								// do not break
								// if the input is less than the size of the buffer,
								// will then use n > 0, n < len(b), and return EOF
							} else {
								fmt.Fprintln(os.Stderr, fmt.Errorf("error reading from resource at uri %q: %w", inputURI, errRead).Error())
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
								fmt.Fprintln(os.Stderr, fmt.Errorf("error writing to resource at uri %q: %w", outputURI, err).Error())
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

			if strings.HasPrefix(outputURI, "s3://") {
				i := strings.Index(outputPath, "/")
				if i == -1 {
					return fmt.Errorf("path missing bucket: %w", err)
				}
				errUpload := grw.UploadS3Object(&grw.UploadS3ObjectInput{
					ACL:    outputACL,
					Bucket: outputPath[0:i],
					Client: s3Client,
					Key:    outputPath[i+1:],
					Object: outputBuffer,
				})
				if errUpload != nil {
					return fmt.Errorf("error uploading to AWS S3 at uri %q: %w", outputURI, errUpload)
				}
			}

			if outputSFTPClient != nil {
				err = outputSFTPClient.Close()
				if err != nil {
					return fmt.Errorf("error closing SFTP client for output: %w", err)
				}
			}
			if outputSSHClient != nil {
				err = outputSSHClient.Close()
				if err != nil {
					return fmt.Errorf("error closing SSH client for output: %w", err)
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

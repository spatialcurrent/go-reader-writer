// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// grw is the command line version of go-reader-writer (GRW) that is used for reading/writing to multiple compression/archive formats.
//
//
//  Usage: grw [-input_uri INPUT_URI] [-input_compression [bzip2|gzip|snappy|zip|none]] [-output_uri OUTPUT_URI] [-output_compression [bzip2|gzip|snappy|zip|none]] [-verbose] [-version]
//  Options:
//  -aws_access_key_id string
//        Defaults to value of environment variable AWS_ACCESS_KEY_ID
//  -aws_default_region string
//        Defaults to value of environment variable AWS_DEFAULT_REGION.
//  -aws_secret_access_key string
//        Defaults to value of environment variable AWS_SECRET_ACCESS_KEY.
//  -aws_session_token string
//        Defaults to value of environment variable AWS_SESSION_TOKEN.
//  -help
//        Print help
//  -input_buffer_size int
//        the input reader buffer size (default 4096)
//  -input_compression string
//        Stream input compression algorithm for nodes, using: bzip2, gzip, snappy, zip, or none.
//  -input_uri string
//        "stdin" or uri to input file (default "stdin")
//  -output_append
//        append output to resource
//  -output_buffer_size int
//        the output writer buffer size (default 4096)
//  -output_compression string
//        Stream input compression algorithm for nodes, using: bzip2, gzip, snappy, zip, or none.
//  -output_uri string
//        "stdout" or uri to output resource (default "stdout")
//  -version
//        Prints version to stdout
//
//
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

import (
	"github.com/spatialcurrent/go-reader-writer/grw"
)

func connect_to_aws(aws_access_key_id string, aws_secret_access_key string, aws_session_token string, aws_region string) *session.Session {
	aws_session := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, aws_session_token),
			MaxRetries:  aws.Int(3),
			Region:      aws.String(aws_region),
		},
	}))
	return aws_session
}

func closeReaderAndWriter(inputReader grw.ByteReadCloser, outputWriter grw.ByteWriteCloser, brokenPipe bool) (error, error) {
	errorReader := inputReader.Close()
	if errorReader != nil {
		fmt.Fprintf(os.Stderr, errors.Wrap(errorReader, "error closing input resource").Error())
	}

	if brokenPipe {
		// close without flushing to output writer
		errorWriter := outputWriter.CloseFile()
		if errorWriter != nil {
			fmt.Fprintf(os.Stderr, errors.Wrap(errorWriter, "error closing output resource").Error())
		}
		return errorReader, errorWriter
	}

	errorWriter := outputWriter.Close()
	if errorWriter != nil {
		fmt.Fprintf(os.Stderr, errors.Wrap(errorWriter, "error closing output resource").Error())
	}

	return errorReader, errorWriter
}

func main() {

	start := time.Now()

	var aws_default_region string
	var aws_access_key_id string
	var aws_secret_access_key string
	var aws_session_token string

	var input_uri string
	var input_compression string
	var input_buffer_size int

	var outputUri string
	var output_compression string
	var output_append bool
	var output_buffer_size int

	var verbose bool
	var version bool
	var help bool

	flag.StringVar(&aws_default_region, "aws_default_region", "", "Defaults to value of environment variable AWS_DEFAULT_REGION.")
	flag.StringVar(&aws_access_key_id, "aws_access_key_id", "", "Defaults to value of environment variable AWS_ACCESS_KEY_ID")
	flag.StringVar(&aws_secret_access_key, "aws_secret_access_key", "", "Defaults to value of environment variable AWS_SECRET_ACCESS_KEY.")
	flag.StringVar(&aws_session_token, "aws_session_token", "", "Defaults to value of environment variable AWS_SESSION_TOKEN.")

	flag.StringVar(&input_uri, "input_uri", "stdin", "\"stdin\" or uri to input file")
	flag.StringVar(&input_compression, "input_compression", "", "Stream input compression algorithm for nodes, using: bzip2, gzip, snappy, zip, or none.")
	flag.IntVar(&input_buffer_size, "input_buffer_size", 4096, "the input reader buffer size") // default from https://golang.org/src/bufio/bufio.go

	flag.IntVar(&output_buffer_size, "output_buffer_size", 4096, "the output writer buffer size")

	flag.StringVar(&outputUri, "output_uri", "stdout", "\"stdout\" or uri to output resource")
	flag.StringVar(&output_compression, "output_compression", "", "Stream input compression algorithm for nodes, using: bzip2, gzip, snappy, zip, or none.")
	flag.BoolVar(&output_append, "output_append", false, "append output to resource")

	flag.BoolVar(&version, "version", false, "Prints version to stdout")
	flag.BoolVar(&help, "help", false, "Print help")

	flag.Parse()

	if len(aws_default_region) == 0 {
		aws_default_region = os.Getenv("AWS_DEFAULT_REGION")
	}
	if len(aws_access_key_id) == 0 {
		aws_access_key_id = os.Getenv("AWS_ACCESS_KEY_ID")
	}
	if len(aws_secret_access_key) == 0 {
		aws_secret_access_key = os.Getenv("AWS_SECRET_ACCESS_KEY")
	}
	if len(aws_session_token) == 0 {
		aws_session_token = os.Getenv("AWS_SESSION_TOKEN")
	}

	if help {
		fmt.Println("Usage: grw [-input_uri INPUT_URI] [-input_compression [bzip2|gzip|snappy|zip|none]] [-output_uri OUTPUT_URI] [-output_compression [bzip2|gzip|snappy|zip|none]] [-verbose] [-version]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	} else if len(os.Args) == 1 {
		fmt.Println("Error: Provided no arguments.")
		fmt.Println("Run \"grw -help\" for more information.")
		os.Exit(0)
	} else if len(os.Args) == 2 && os.Args[1] == "help" {
		fmt.Println("Usage: grw [-input_uri INPUT_URI] [-input_compression [bzip2|gzip|snappy|zip|none]] [-output_uri OUTPUT_URI] [-output_compression [bzip2|gzip|snappy|zip|none]] [-verbose] [-version]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if version {
		fmt.Println(grw.Version)
		os.Exit(0)
	}

	var aws_session *session.Session
	var s3_client *s3.S3

	if strings.HasPrefix(input_uri, "s3://") || strings.HasPrefix(outputUri, "s3://") {
		aws_session = connect_to_aws(aws_access_key_id, aws_secret_access_key, aws_session_token, aws_default_region)
		s3_client = s3.New(aws_session)
	}

	inputReader, _, err := grw.ReadFromResource(input_uri, input_compression, input_buffer_size, false, s3_client)
	if err != nil {
		fmt.Fprintf(os.Stderr, errors.Wrap(err, "error opening resource at uri "+input_uri).Error())
		os.Exit(1)
	}

	var outputWriter grw.ByteWriteCloser
	var outputBuffer *bytes.Buffer
	if strings.HasPrefix(outputUri, "s3://") {
		outputWriter, outputBuffer, err = grw.WriteBytes(output_compression)
		if err != nil {
			fmt.Fprintf(os.Stderr, errors.Wrap(err, "error opening bytes buffer for "+outputUri).Error())
			os.Exit(1)
		}
	} else {
		outputWriter, err = grw.WriteToResource(outputUri, output_compression, output_append, s3_client)
		if err != nil {
			fmt.Fprintf(os.Stderr, errors.Wrap(err, "error opening resource at uri "+outputUri).Error())
			os.Exit(1)
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

			b := make([]byte, output_buffer_size)
			n, err := inputReader.Read(b)
			if err != nil {
				if err == io.EOF {
					eof = true
				} else {
					fmt.Fprintf(os.Stderr, errors.Wrap(err, "error reading from resource at uri "+input_uri).Error())
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
				fmt.Fprintf(os.Stderr, errors.Wrap(err, "error writing to resource at uri "+outputUri).Error())
			}

		}
		wg.Done()
	}()

	wg.Wait() // wait until done writing or received signal for graceful shutdown

	errorReader, errorWriter := closeReaderAndWriter(inputReader, outputWriter, brokenPipe)
	if errorReader != nil || errorWriter != nil {
		os.Exit(1)
	}

	if strings.HasPrefix(outputUri, "s3://") {
		_, outputPath := grw.SplitUri(outputUri)
		i := strings.Index(outputPath, "/")
		if i == -1 {
			fmt.Fprintf(os.Stderr, errors.Wrap(err, "path missing bucket").Error())
			os.Exit(1)
		}
		err := grw.UploadS3Object(outputPath[0:i], outputPath[i+1:], outputBuffer, s3_client)
		if err != nil {
			fmt.Fprintf(os.Stderr, errors.Wrap(err, "error uploading to AWS S3 at uri "+outputUri).Error())
			os.Exit(1)
		}
	}

	elapsed := time.Since(start)
	if verbose && !brokenPipe {
		fmt.Println("Done in " + elapsed.String())
	}

}

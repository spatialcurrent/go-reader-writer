// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// +build !js

// Package grw provides the interfaces, embedded structs, and implementing code
// for normalizing the reading/writing of a stream of bytes from archive/compressed files.
// This package supports the bzip2, gzip, snappy, and zip archive/compression algorithms.  No compression can be identified as "none" or a blank string.
// This package is used by the go-stream package.
//  - https://godoc.org/github.com/spatialcurrent/go-stream/stream
//
// Usage
//
// You can import reader as a package into your own Go project or use the command line interface.
//
//  import (
//    "github.com/spatialcurrent/go-reader-writer/grw"
//  )
//
//  r, err := grw.ReadFromFilePath("data-for-2018.sz", "snappy", false, 4096)
//  if err != nil {
//    panic(err)
//  }
//  for {
//    b, err := input_reader.ReadBytes([]byte("\n")[0])
//    if err != nil {
//      if err != io.EOF {
//        fmt.Println(errors.Wrap(err, "Error reading bytes from file"))
//        os.Exit(1)
//      }
//    }
//    if len(b) > 0 {
//      fmt.Println(string(b))
//    }
//    if err != nil && err == io.EOF {
//      break
//    }
//  }
//
//
// See the github.com/go-reader/cmd/go-reader package for a command line tool for testing DFL expressions.
//
//  - https://godoc.org/github.com/spatialcurrent/go-reader-writer/grw
//
// Projects
// go-reader-writer is used by the railgun project and go-stream
//
//  - https://github.com/spatialcurrent/railgun
//  - https://godoc.org/pkg/github.com/spatialcurrent/go-stream/stream
//
package grw

import (
	"errors"
)

const (
	AlgorithmBzip2  = "bzip2"  // bzip2
	AlgorithmFlate  = "flate"  // flate aka DEFLATE
	AlgorithmGzip   = "gzip"   // gzip
	AlgorithmNone   = "none"   // no compressions
	AlgorithmSnappy = "snappy" // snappy
	AlgorithmZip    = "zip"    // zip archive
	AlgorithmZlib   = "zlib"   // zlib

	SchemeFile  = "file"
	SchemeFtp   = "ftp"
	SchemeHTTP  = "http"
	SchemeHTTPS = "https"
	SchemeS3    = "s3"
)

var (
	Algorithms = []string{
		AlgorithmBzip2,
		AlgorithmFlate,
		AlgorithmGzip,
		AlgorithmNone,
		AlgorithmSnappy,
		AlgorithmZip,
		AlgorithmZlib,
	}
)

var (
	Schemes = []string{
		SchemeFile,
		SchemeFtp,
		SchemeHTTP,
		SchemeHTTPS,
		SchemeS3,
	}
)

var (
	ErrPathMissing = errors.New("path is missing")
)

var (
	DefaultBufferSize = 4096
)

var (
	NoDict = []byte{} // no dictionary
)

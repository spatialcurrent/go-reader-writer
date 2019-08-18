// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// +build js

package grw

import (
	"errors"
)

const (
	AlgorithmBzip2  = "bzip2"  // bzip2
	AlgorithmGzip   = "gzip"   // gzip
	AlgorithmNone   = "none"   // No compression
	AlgorithmSnappy = "snappy" // Snappy compression
	AlgorithmZip    = "zip"    // Zip archive

	SchemeFile  = "file"
	SchemeHTTP  = "http"
	SchemeHTTPS = "https"
	SchemeS3    = "s3"
)

var (
	Algorithms = []string{
		AlgorithmBzip2,
		AlgorithmGzip,
		AlgorithmNone,
		AlgorithmSnappy,
		AlgorithmZip,
	}

	Schemes = []string{
		SchemeFile,
		SchemeHTTP,
		SchemeHTTPS,
		SchemeS3,
	}

	ErrPathMissing = errors.New("path is missing")
)

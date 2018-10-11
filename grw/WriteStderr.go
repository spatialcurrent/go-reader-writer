// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"compress/gzip"
	"os"
)

import (
	"github.com/golang/snappy"
)

// WriteStderr returns a ByteWriteCloser for stderr with the given compression.
// alg may be "bzip2", "gzip", "snappy", "zip", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//  - https://golang.org/pkg/archive/zip/
//
func WriteStderr(alg string) (ByteWriteCloser, error) {
	switch alg {
	case "bzip2":
		return nil, &ErrWriterNotImplemented{Algorithm: alg}
	case "gzip":
		gw := gzip.NewWriter(os.Stdout)
		return &Writer{Writer: bufio.NewWriter(gw), Closer: gw}, nil
	case "snappy":
		sw := snappy.NewWriter(os.Stdout)
		return &Writer{Writer: bufio.NewWriter(sw), Closer: sw}, nil
	case "zip":
		return nil, &ErrWriterNotImplemented{Algorithm: alg}
	case "none", "":
		return &Writer{Writer: bufio.NewWriter(os.Stdout)}, nil
	}

	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

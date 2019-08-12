// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
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

// WriteStdout returns a ByteWriteCloser for stderr with the given compression.
// alg may be "bzip2", "gzip", "snappy", "zip", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//  - https://golang.org/pkg/archive/zip/
//
func WriteStdout(alg string) (ByteWriteCloser, error) {
	switch alg {
	case AlgorithmBzip2:
		return nil, &ErrWriterNotImplemented{Algorithm: alg}
	case AlgorithmGzip:
		gw := gzip.NewWriter(os.Stdout)
		return NewWriterWithCloser(bufio.NewWriter(gw), gw), nil
	case AlgorithmSnappy:
		sw := snappy.NewWriter(os.Stdout)
		return NewWriterWithCloser(bufio.NewWriter(sw), sw), nil
	case AlgorithmZip:
		return nil, &ErrWriterNotImplemented{Algorithm: alg}
	case AlgorithmNone, "":
		return NewWriter(bufio.NewWriter(os.Stdout)), nil
	}

	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

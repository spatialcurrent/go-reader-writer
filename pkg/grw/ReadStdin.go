// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"compress/bzip2"
	"compress/gzip"
	"io"
	"os"
)

import (
	"github.com/golang/snappy"
	"github.com/pkg/errors"
)

// ReadStdin returns a ByteReadCloser for a file with a given compression.
// alg may be "bzip2", "gzip", "snappy", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func ReadStdin(alg string) (ByteReadCloser, error) {
	switch alg {
	case "bzip2":
		return &Reader{Reader: bufio.NewReader(bzip2.NewReader(os.Stdin))}, nil
	case "gzip":
		gr, err := gzip.NewReader(bufio.NewReader(os.Stdin))
		if err != nil {
			return nil, errors.Wrap(err, "Error creating gzip reader for stdin")
		}
		return &Reader{Reader: bufio.NewReader(gr), Closer: &Closer{Closers: []io.Closer{gr}}}, nil
	case "snappy":
		return &Reader{Reader: bufio.NewReader(snappy.NewReader(bufio.NewReader(os.Stdin)))}, nil
	case "none", "":
		return &Reader{Reader: bufio.NewReader(os.Stdin)}, nil
	}
	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

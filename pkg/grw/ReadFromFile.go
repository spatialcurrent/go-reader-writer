// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"io"
	"os"
)

import (
	"github.com/pkg/errors"
)

// ReadFromFile returns a ByteReader for a file with a given compression.
// alg may be "snappy", "gzip", or "none."
//
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func ReadFromFile(file *os.File, alg string, bufferSize int) (ByteReadCloser, error) {
	switch alg {
	case AlgorithmBzip2, AlgorithmGzip, AlgorithmSnappy, AlgorithmNone, "":
		r, closers, err := WrapReader(file, []io.Closer{file}, alg, bufferSize)
		if err != nil {
			return nil, errors.Wrapf(err, "error wrapping reader for file at path %q", file.Name())
		}
		return &Reader{Reader: r, Closers: closers}, nil
	case AlgorithmZip:
		brc, err := ReadZipFile(file.Name())
		if err != nil {
			return nil, errors.Wrapf(err, "error creating reader for zip file at path %q", file.Name())
		}
		return brc, nil
	}
	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

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

	"github.com/pkg/errors"
)

// ReadFromFile opens a ByteReadCloser for the given path, compression algorithm, and buffer size.
// ReadFromFile returns the ByteReadCloser and error, if any.
//
// ReadFromFile returns an error if the compression algorithm is invalid.
//
func ReadFromFile(file *os.File, alg string, bufferSize int) (ByteReadCloser, error) {
	switch alg {
	case AlgorithmBzip2, AlgorithmGzip, AlgorithmSnappy, AlgorithmNone, "":
		r, closers, err := WrapReader(file, []io.Closer{file}, alg, bufferSize)
		if err != nil {
			return nil, errors.Wrapf(err, "error wrapping reader for file at path %q", file.Name())
		}
		return &Reader{Reader: r, Closer: &Closer{Closers: closers}}, nil
	case AlgorithmZip:
		brc, err := ReadZipFile(file.Name())
		if err != nil {
			return nil, errors.Wrapf(err, "error creating reader for zip file at path %q", file.Name())
		}
		return brc, nil
	}
	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

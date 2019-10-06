// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"os"

	"github.com/pkg/errors"
)

type ReadFromFileInput struct {
	File       *os.File // file to read from
	Alg        string   // compression algorithm
	Dict       []byte   // compression dictionary
	BufferSize int      // input reader buffer size
}

// ReadFromFile opens a ByteReadCloser for the given path, compression algorithm, and buffer size.
// ReadFromFile returns the ByteReadCloser and error, if any.
//
// ReadFromFile returns an error if the compression algorithm is invalid.
//
func ReadFromFile(input *ReadFromFileInput) (*Reader, error) {
	switch input.Alg {
	case AlgorithmBzip2, AlgorithmGzip, AlgorithmSnappy, AlgorithmNone, "":
		r, err := WrapReader(input.File, input.Alg, input.Dict, input.BufferSize)
		if err != nil {
			return nil, errors.Wrapf(err, "error wrapping reader for file at path %q", input.File.Name())
		}
		return &Reader{Reader: r}, nil
	case AlgorithmZip:
		brc, err := ReadZipFile(input.File.Name())
		if err != nil {
			return nil, errors.Wrapf(err, "error creating reader for zip file at path %q", input.File.Name())
		}
		return brc, nil
	}
	return nil, &ErrUnknownAlgorithm{Algorithm: input.Alg}
}

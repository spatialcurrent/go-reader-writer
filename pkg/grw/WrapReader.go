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
)

import (
	"github.com/golang/snappy"
	"github.com/pkg/errors"
)

func WrapReader(r io.Reader, closers []io.Closer, alg string, bufferSize int) (ByteReader, []io.Closer, error) {
	var br ByteReader
	switch alg {
	case AlgorithmBzip2:
		br = bufio.NewReaderSize(bzip2.NewReader(bufio.NewReaderSize(r, bufferSize)), bufferSize)
	case AlgorithmGzip:
		gr, err := gzip.NewReader(bufio.NewReaderSize(r, bufferSize))
		if err != nil {
			return nil, nil, errors.Wrap(err, "error creating gzip reader for reader")
		}
		br = bufio.NewReaderSize(gr, bufferSize)
		closers = append([]io.Closer{gr}, closers...)
	case AlgorithmSnappy:
		br = bufio.NewReaderSize(snappy.NewReader(bufio.NewReaderSize(r, bufferSize)), bufferSize)
	case AlgorithmNone, "":
		br = bufio.NewReaderSize(r, bufferSize)
	}
	return br, closers, nil
}

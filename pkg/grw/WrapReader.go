// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/bzip2"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/flate"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/gzip"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/snappy"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/zlib"
	"github.com/spatialcurrent/go-reader-writer/pkg/io"
)

func WrapReader(r io.Reader, alg string, dict []byte, bufferSize int) (io.ByteReader, error) {
	var br io.ByteReader
	switch alg {
	case AlgorithmBzip2:
		br = bufio.NewReader(bzip2.NewReader(bufio.NewReaderSize(r, bufferSize)))
	case AlgorithmFlate:
		if len(dict) > 0 {
			br = bufio.NewReader(flate.NewReaderDict(bufio.NewReaderSize(r, bufferSize), dict))
		} else {
			br = bufio.NewReader(flate.NewReader(bufio.NewReaderSize(r, bufferSize)))
		}
	case AlgorithmGzip:
		gr, err := gzip.NewReader(bufio.NewReaderSize(r, bufferSize))
		if err != nil {
			return nil, errors.Wrap(err, "error creating gzip reader for reader")
		}
		br = bufio.NewReader(gr)
	case AlgorithmSnappy:
		br = bufio.NewReader(snappy.NewReader(bufio.NewReaderSize(r, bufferSize)))
	case AlgorithmZlib:
		if len(dict) > 0 {
			zr, err := zlib.NewReaderDict(bufio.NewReaderSize(r, bufferSize), dict)
			if err != nil {
				return nil, errors.Wrap(err, "error creating zlib reader with dictionary for reader")
			}
			br = bufio.NewReader(zr)
		} else {
			zr, err := zlib.NewReader(bufio.NewReaderSize(r, bufferSize))
			if err != nil {
				return nil, errors.Wrap(err, "error creating zlib reader for reader")
			}
			br = bufio.NewReader(zr)
		}
	case AlgorithmNone, "":
		br = bufio.NewReaderSize(r, bufferSize)
	}
	return br, nil
}

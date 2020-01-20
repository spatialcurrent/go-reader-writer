// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/bzip2"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/flate"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/gzip"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/snappy"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/zlib"
	"github.com/spatialcurrent/go-reader-writer/pkg/io"
)

func WrapReader(r io.Reader, alg string, dict []byte, bufferSize int) (io.Reader, error) {
	switch alg {
	case AlgorithmBzip2:
		return bufio.NewReader(bzip2.NewReader(bufio.NewReaderSize(r, bufferSize))), nil
	case AlgorithmFlate:
		if len(dict) > 0 {
			return bufio.NewReader(flate.NewReaderDict(bufio.NewReaderSize(r, bufferSize), dict)), nil
		}
		return bufio.NewReader(flate.NewReader(bufio.NewReaderSize(r, bufferSize))), nil
	case AlgorithmGzip:
		gr, err := gzip.NewReader(bufio.NewReaderSize(r, bufferSize))
		if err != nil {
			return nil, fmt.Errorf("error creating gzip reader for reader: %w", err)
		}
		return bufio.NewReader(gr), nil
	case AlgorithmSnappy:
		return bufio.NewReader(snappy.NewReader(bufio.NewReaderSize(r, bufferSize))), nil
	case AlgorithmZlib:
		if len(dict) > 0 {
			zr, err := zlib.NewReaderDict(bufio.NewReaderSize(r, bufferSize), dict)
			if err != nil {
				return nil, fmt.Errorf("error creating zlib reader with dictionary for reader: %w", err)
			}
			return bufio.NewReader(zr), nil
		} else {
			zr, err := zlib.NewReader(bufio.NewReaderSize(r, bufferSize))
			if err != nil {
				return nil, fmt.Errorf("error creating zlib reader for reader: %w", err)
			}
			return bufio.NewReader(zr), nil
		}
	case AlgorithmNone, "":
		// if buffer size is zero, then don't wrap with bufio
		if bufferSize == 0 {
			return r, nil
		}
		return bufio.NewReaderSize(r, bufferSize), nil
	}
	return nil, fmt.Errorf("unsupported compression algorith %q", alg)
}

// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"

	pkgalg "github.com/spatialcurrent/go-reader-writer/pkg/alg"
	"github.com/spatialcurrent/go-reader-writer/pkg/archive/zip"
	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/bzip2"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/flate"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/gzip"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/snappy"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/zlib"
	"github.com/spatialcurrent/go-reader-writer/pkg/io"
)

func WrapReader(r io.ReadCloser, alg string, dict []byte, bufferSize int) (io.ReadCloser, error) {
	switch alg {
	case pkgalg.AlgorithmBzip2:
		return bufio.NewReader(bzip2.NewReader(bufio.NewReaderSize(r, bufferSize))), nil
	case pkgalg.AlgorithmFlate:
		if len(dict) > 0 {
			return bufio.NewReader(flate.NewReaderDict(bufio.NewReaderSize(r, bufferSize), dict)), nil
		}
		return flate.NewReader(bufio.NewReaderSize(r, bufferSize)), nil
	case pkgalg.AlgorithmGzip:
		gr, err := gzip.NewReader(bufio.NewReaderSize(r, bufferSize))
		if err != nil {
			return nil, fmt.Errorf("error creating gzip reader for reader: %w", err)
		}
		return gr, nil
	case pkgalg.AlgorithmSnappy:
		return snappy.NewReader(bufio.NewReaderSize(r, bufferSize)), nil
	case pkgalg.AlgorithmZlib:
		if len(dict) > 0 {
			zr, err := zlib.NewReaderDict(bufio.NewReaderSize(r, bufferSize), dict)
			if err != nil {
				return nil, fmt.Errorf("error creating zlib reader with dictionary for reader: %w", err)
			}
			return zr, nil
		} else {
			zr, err := zlib.NewReader(bufio.NewReaderSize(r, bufferSize))
			if err != nil {
				return nil, fmt.Errorf("error creating zlib reader for reader: %w", err)
			}
			return zr, nil
		}
	case pkgalg.AlgorithmZip:
		b, err := io.ReadAll(r)
		if err != nil {
			return nil, fmt.Errorf("error creating ZIP reader for reader: %w", err)
		}
		zr, err := zip.ReadBytes(b)
		if err != nil {
			return nil, fmt.Errorf("error creating ZIP reader for reader: %w", err)
		}
		return zr, nil
	case pkgalg.AlgorithmNone, "":
		// if buffer size is zero, then don't wrap with bufio
		if bufferSize == 0 {
			return r, nil
		}
		return bufio.NewReaderSize(r, bufferSize), nil
	}

	return nil, &pkgalg.ErrUnknownAlgorithm{Algorithm: alg}
}

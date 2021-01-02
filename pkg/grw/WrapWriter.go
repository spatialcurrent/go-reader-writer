// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
	"io"

	pkgalg "github.com/spatialcurrent/go-reader-writer/pkg/alg"
	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/flate"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/gzip"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/snappy"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/zlib"
)

// WrapWriter wraps the given writer with a buffer and the given compression.
// alg is the algorithm.  dict is the initial dictionary (if the algorithm uses one).
//
//  - https://pkg.go.dev/pkg/archive/zip/
//  - https://pkg.go.dev/pkg/compress/bzip2/
//  - https://pkg.go.dev/pkg/compress/flate/
//  - https://pkg.go.dev/pkg/compress/gzip/
//  - https://pkg.go.dev/pkg/compress/zlib/
//  - https://pkg.go.dev/pkg/github.com/golang/snappy
//  - https://pkg.go.dev/pkg/github.com/go-reader-writer/pkg/bufio
//
func WrapWriter(w io.WriteCloser, alg string, dict []byte, bufferSize int) (io.WriteCloser, error) {
	if bufferSize < 0 {
		return nil, fmt.Errorf("error wrapping writer: invalid buffer size of %d", bufferSize)
	}
	switch alg {
	case pkgalg.AlgorithmBzip2:
		return nil, &ErrWriterNotImplemented{Algorithm: alg}
	case pkgalg.AlgorithmFlate:
		if len(dict) > 0 {
			fw, err := flate.NewWriterDict(bufio.NewWriter(w), flate.DefaultCompression, dict)
			if err != nil {
				return nil, fmt.Errorf("error wrapping writer using compression %q with dictionary %q: %w", alg, string(dict), err)
			}
			return fw, nil
		}
		fw, err := flate.NewWriter(bufio.NewWriter(w))
		if err != nil {
			return nil, fmt.Errorf("error wrapping writer using compression %q: %w", alg, err)
		}
		return fw, nil
	case pkgalg.AlgorithmGzip:
		return gzip.NewWriter(bufio.NewWriter(w)), nil
	case pkgalg.AlgorithmSnappy:
		return snappy.NewBufferedWriter(bufio.NewWriter(w)), nil
	case pkgalg.AlgorithmZip:
		return nil, &ErrWriterNotImplemented{Algorithm: alg}
	case pkgalg.AlgorithmZlib:
		if len(dict) > 0 {
			zw, err := zlib.NewWriterDict(bufio.NewWriter(w), dict)
			if err != nil {
				return nil, fmt.Errorf("error wrapping writer using compression %q with dictionary %q: %w", alg, string(dict), err)
			}
			return zw, nil
		}
		return zlib.NewWriter(bufio.NewWriter(w)), nil
	case pkgalg.AlgorithmNone:
		if bufferSize > 0 {
			return bufio.NewWriter(w), nil
		}
		return w, nil
	}
	return nil, &pkgalg.ErrUnknownAlgorithm{Algorithm: alg}
}

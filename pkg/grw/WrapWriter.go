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
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/flate"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/gzip"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/snappy"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/zlib"
	"github.com/spatialcurrent/go-reader-writer/pkg/io"
)

// WrapWriter wraps the given writer with a buffer and the given compression.
// alg is the algorithm.  dict is the initial dictionary (if the algorithm uses one).
//
//  - https://godoc.org/pkg/archive/zip/
//  - https://godoc.org/pkg/compress/bzip2/
//  - https://godoc.org/pkg/compress/flate/
//  - https://godoc.org/pkg/compress/gzip/
//  - https://godoc.org/pkg/compress/zlib/
//  - https://godoc.org/pkg/github.com/golang/snappy
//  - https://godoc.org/pkg/github.com/go-reader-writer/pkg/bufio
//
func WrapWriter(w io.Writer, alg string, dict []byte) (*Writer, error) {
	switch alg {
	case AlgorithmBzip2:
		return nil, &ErrWriterNotImplemented{Algorithm: alg}
	case AlgorithmFlate:
		if len(dict) > 0 {
			fw, err := flate.NewWriterDict(w, flate.DefaultCompression, dict)
			if err != nil {
				return nil, errors.Wrapf(err, "error wrapping writer using compression %q with dictionary %q", alg, string(dict))
			}
			return NewWriter(bufio.NewWriter(fw)), nil
		}
		fw, err := flate.NewWriter(w, flate.DefaultCompression)
		if err != nil {
			return nil, errors.Wrapf(err, "error wrapping writer using compression %q", alg)
		}
		return NewWriter(bufio.NewWriter(fw)), nil
	case AlgorithmGzip:
		return NewWriter(bufio.NewWriter(gzip.NewWriter(w))), nil
	case AlgorithmSnappy:
		return NewWriter(bufio.NewWriter(snappy.NewBufferedWriter(w))), nil
	case AlgorithmZip:
		return nil, &ErrWriterNotImplemented{Algorithm: alg}
	case AlgorithmZlib:
		if len(dict) > 0 {
			zw, err := zlib.NewWriterDict(w, dict)
			if err != nil {
				return nil, errors.Wrapf(err, "error wrapping writer using compression %q with dictionary %q", alg, string(dict))
			}
			return NewWriter(bufio.NewWriter(zw)), nil
		}
		return NewWriter(bufio.NewWriter(zlib.NewWriter(w))), nil
	case AlgorithmNone, "":
		return NewWriter(bufio.NewWriter(w)), nil
	}
	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

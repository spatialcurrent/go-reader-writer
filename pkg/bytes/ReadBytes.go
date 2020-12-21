// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package bytes

import (
	"io"
	"io/ioutil"

	pkgalg "github.com/spatialcurrent/go-reader-writer/pkg/alg"
)

// ReadBytes returns a ByteReader for a byte array with a given compression.
// alg may be "flate", "gzip", "none", "snappy", or "zlib".
//
//  - https://pkg.go.dev/pkg/compress/flate/
//  - https://pkg.go.dev/pkg/compress/gzip/
//  - https://pkg.go.dev/pkg/compress/zlib/
//  - https://pkg.go.dev/github.com/golang/snappy
//
func ReadBytes(b []byte, alg string, dict []byte) (io.ReadCloser, error) {
	switch alg {
	case pkgalg.AlgorithmFlate:
		return ReadFlateBytes(b, dict), nil
	case pkgalg.AlgorithmGzip:
		return ReadGzipBytes(b, true)
	case pkgalg.AlgorithmSnappy:
		return ioutil.NopCloser(ReadSnappyBytes(b)), nil
	case pkgalg.AlgorithmZlib:
		return ReadZlibBytes(b, dict)
	case "none", "":
		return ioutil.NopCloser(ReadPlainBytes(b)), nil
	}
	return nil, &pkgalg.ErrUnknownAlgorithm{Algorithm: alg}
}

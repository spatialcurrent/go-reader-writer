// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

// ReadBytes returns a ByteReader for a byte array with a given compression.
// alg may be "snappy", "gzip", or "none."
//
//  - https://golang.org/pkg/compress/flate/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func ReadBytes(b []byte, alg string, dict []byte) (*Reader, error) {
	switch alg {
	case AlgorithmFlate:
		return ReadFlateBytes(b, dict), nil
	case AlgorithmGzip:
		return ReadGzipBytes(b)
	case AlgorithmSnappy:
		return ReadSnappyBytes(b)
	case AlgorithmZlib:
		return ReadZlibBytes(b, dict)
	case "none", "":
		return ReadMemoryBytes(b), nil
	}
	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

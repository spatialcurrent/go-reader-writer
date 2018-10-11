// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

// ReadBytes returns a ByteReader for a byte array with a given compression.
// alg may be "snappy", "gzip", or "none."
//
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func ReadBytes(b []byte, alg string) (ByteReadCloser, error) {
	switch alg {
	case "snappy":
		return ReadSnappyBytes(b)
	case "gzip":
		return ReadGzipBytes(b)
	case "none", "":
		return ReadMemoryBytes(b)
	}
	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bytes"
)

// WriteBytes returns a ByteReader for a byte array with a given compression.
// alg may be "snappy", "gzip", or "none."
//
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func WriteBytes(alg string) (ByteWriteCloser, *bytes.Buffer, error) {
	switch alg {
	case "bzip2":
		return nil, nil, &ErrWriterNotImplemented{Algorithm: alg}
	case "gzip":
		return WriteGzipBytes()
	case "snappy":
		return WriteSnappyBytes()
	case "zip":
		return nil, nil, &ErrWriterNotImplemented{Algorithm: alg}
	case "none", "":
		return WriteMemoryBytes()
	}
	return nil, new(bytes.Buffer), &ErrUnknownAlgorithm{Algorithm: alg}
}

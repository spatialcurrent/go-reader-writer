// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
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
//  - https://golang.org/pkg/compress/flate/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func WriteBytes(alg string, dict []byte) (*Writer, *bytes.Buffer, error) {
	switch alg {
	case AlgorithmBzip2:
		return nil, nil, &ErrWriterNotImplemented{Algorithm: alg}
	case AlgorithmGzip:
		writer, buffer := WriteGzipBytes()
		return writer, buffer, nil
	case AlgorithmFlate:
		return WriteFlateBytes(dict)
	case AlgorithmSnappy:
		writer, buffer := WriteSnappyBytes()
		return writer, buffer, nil
	case AlgorithmZip:
		return nil, nil, &ErrWriterNotImplemented{Algorithm: alg}
	case AlgorithmNone, "":
		writer, buffer := WriteMemoryBytes()
		return writer, buffer, nil
	}
	return nil, nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

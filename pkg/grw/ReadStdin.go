// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
	"os"
)

// ReadStdin returns a ByteReadCloser for a file with a given compression.
// alg may be "bzip2", "flate", "gzip", "snappy", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/flate/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func ReadStdin(alg string, dict []byte, bufferSize int) (*Reader, error) {

	r, err := WrapReader(os.Stdin, alg, dict, bufferSize)
	if err != nil {
		return nil, fmt.Errorf("error wrapping reader for stdin: %w", err)
	}
	return &Reader{Reader: r}, nil
}

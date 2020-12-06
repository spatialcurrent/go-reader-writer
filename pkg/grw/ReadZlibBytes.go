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
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/zlib"
)

// ReadZlibBytes returns a reader for reading zlib bytes from an input slice.
// Wraps the "compress/zlib" package.
//
//  - https://golang.org/pkg/compress/zlib/
//
func ReadZlibBytes(b []byte, dict []byte) (*Reader, error) {
	zr, err := zlib.ReadBytes(b, dict)
	if err != nil {
		return nil, fmt.Errorf("error creating zlib reader for memory block: %w", err)
	}
	return &Reader{Reader: bufio.NewReader(zr)}, nil
}

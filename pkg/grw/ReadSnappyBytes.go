// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"

	"github.com/spatialcurrent/go-reader-writer/pkg/compress/snappy"
)

// SnappyBytes returns a reader for an input of snappy-compressed bytes, and an error if any.
//
//  - https://godoc.org/github.com/golang/snappy
//
func ReadSnappyBytes(b []byte) (*Reader, error) {
	return &Reader{Reader: bufio.NewReader(snappy.ReadBytes(b))}, nil
}

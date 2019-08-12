// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"bytes"
	"github.com/golang/snappy"
)

// SnappyBytes returns a reader for an input of snappy-compressed bytes, and an error if any.
//
//  - https://godoc.org/github.com/golang/snappy
//
func ReadSnappyBytes(b []byte) (ByteReadCloser, error) {
	return &Reader{Reader: bufio.NewReader(snappy.NewReader(bytes.NewReader(b)))}, nil
}

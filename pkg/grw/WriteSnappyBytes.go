// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bytes"

	"github.com/golang/snappy"
)

// WriteSnappyBytes returns a reader for reading the bytes from an input array, and an error if any.
func WriteSnappyBytes() (ByteWriteCloser, *bytes.Buffer) {
	buf := new(bytes.Buffer)
	sw := snappy.NewBufferedWriter(buf)
	return NewBufferedWriterWithClosers(sw, sw), buf
}

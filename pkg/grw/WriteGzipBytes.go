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
	"compress/gzip"
)

// WriteGzipBytes returns a reader for reading the bytes from an input array, and an error if any.
func WriteGzipBytes() (ByteWriteCloser, Buffer, error) {
	buf := new(bytes.Buffer)
	gw := gzip.NewWriter(buf)
	return NewWriterWithCloser(bufio.NewWriter(gw), gw), buf, nil
}

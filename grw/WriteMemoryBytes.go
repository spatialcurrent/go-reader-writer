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
)

// WriteMemoryBytes returns a reader for reading the bytes from an input array, and an error if any.
func WriteMemoryBytes() (ByteWriteCloser, *bytes.Buffer) {
	buf := new(bytes.Buffer)
	bw := bufio.NewWriter(buf)
	return NewWriter(bw), buf
}

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

// ReadMemoryBytes returns a reader for reading the bytes from an input array, and an error if any.
func ReadMemoryBytes(b []byte) *Cache {
	return NewCache(&Reader{Reader: bufio.NewReader(bytes.NewReader(b))})
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package io

import (
	"io"
)

// Buffer is an interface that supports common buffer methods, including those from the bytes.Buffer struct.
// Buffer extends io.Reader, io.Writer, and io.ByteWriter interfaces.
// It supports other buffer implementations, too.
type Buffer interface {
	io.Reader
	Writer
	io.ByteWriter
	WriteRune(r rune) (n int, err error)
	WriteString(s string) (n int, err error)
	Bytes() []byte
	String() string
	Len() int
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"io"
)

type Buffer interface {
	io.Reader
	io.Writer
	io.ByteWriter
	WriteRune(r rune) (n int, err error)
	WriteString(s string) (n int, err error)
	Bytes() []byte
	String() string
	Len() int
}

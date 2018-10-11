// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"io"
)

// ByteWriteCloser is an interface that extends io.Writer, io.ByteWriter, and io.Closer
type ByteWriteCloser interface {
	ByteWriter
	io.Closer
	CloseFile() error
	WriteString(s string) (n int, err error)
	WriteError(e error) (n int, err error)
}

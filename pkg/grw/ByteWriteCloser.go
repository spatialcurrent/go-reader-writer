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

// ByteWriteCloser is an interface that extends io.Writer, io.ByteWriter, and io.Closer
type ByteWriteCloser interface {
	ByteWriter
	Flusher
	io.Closer
	Lock()
	Unlock()
	FlushSafe() error
	CloseSafe() error
	WriteString(s string) (n int, err error)
	WriteLine(s string) (n int, err error)
	WriteLineSafe(s string) (n int, err error)
	WriteError(e error) (n int, err error)
	WriteErrorSafe(e error) (n int, err error)
}

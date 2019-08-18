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

// ByteReader is an interface that extends io.Reader, io.ByteReader, and adds a range function.
// ByteReader provides functions for reading bytes.
type ByteReadCloser interface {
	ByteReader
	io.Closer
	io.ReaderAt
	ReadAll() ([]byte, error)
	ReadAllAndClose() ([]byte, error)
	ReadFirst() (byte, error)
	ReadRange(start int, end int) ([]byte, error)
}

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

// ByteReadCloser extends ByteReader, io.Closer, and io.ReaderAt interfaces, and provides functions for reading a range of bytes.
type ByteReadCloser interface {
	ByteReader
	io.Closer
	io.ReaderAt
	ReadAll() ([]byte, error)
	ReadAllAndClose() ([]byte, error)
	ReadFirst() (byte, error)
	ReadRange(start int, end int) ([]byte, error)
}

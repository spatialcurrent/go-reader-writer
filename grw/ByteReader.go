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

// ByteReader is an interface that extends io.Reader and io.ByteReader.
// ByteReader provides functions for reading bytes.
type ByteReader interface {
	io.Reader
	io.ByteReader
	ReadBytes(delim byte) ([]byte, error)
}

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

// ByteReader is an interface that supports reading bytes.
// Buffer extends io.Reader io.ByteReader interfaces.
type ByteReader interface {
	io.Reader
	io.ByteReader
	ReadBytes(delim byte) ([]byte, error)
	ReadString(delim byte) (string, error)
}

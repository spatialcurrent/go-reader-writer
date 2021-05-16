// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package bytes

import (
	"io"
)

type ByteReader interface {
	io.Reader
	io.ByteReader
}

type ByteReadScanner interface {
	io.Reader
	io.ByteReader
	io.ByteScanner
}

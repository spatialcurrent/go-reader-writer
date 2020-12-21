// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package flate provides a reader and writer that propagate calls to Flush and Close.
package flate

import (
	"io"
)

// Resetter resets a ReadCloser returned by NewReader or NewReaderDict
// to switch to a new underlying Reader. This permits reusing a ReadCloser
// instead of allocating a new one.
type Resetter interface {
	Reset(r io.Reader, dict []byte) error
}

type ReadResetCloser interface {
	io.Reader
	io.Closer
	Resetter
}

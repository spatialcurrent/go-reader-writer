// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package gzip provides a reader and writer that propagate calls to Flush and Close.
package gzip

import (
	"io"
)

type Resetter interface {
	Reset(r io.Reader) error
}

type ReadResetCloser interface {
	io.Reader
	io.Closer
	Resetter
}

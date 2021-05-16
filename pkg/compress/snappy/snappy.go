// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package snappy provides a reader and writer that propagate calls to Flush and Close.
package snappy

import (
	"io"
)

type Resetter interface {
	Reset(reader io.Reader)
}

type ReadResetter interface {
	io.Reader
	Resetter
}

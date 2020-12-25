// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package io

import (
	"io"
)

// WriteCloser is an interface that extends io.Writer and io.Closer.
type WriteCloser interface {
	Writer
	io.Closer
}

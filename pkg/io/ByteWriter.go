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

// ByteWriter is an interface that extends io.Writer and io.ByteWriter.
// ByteWriter provides functions for writing bytes.
type ByteWriter interface {
	Writer
	io.ByteWriter
}

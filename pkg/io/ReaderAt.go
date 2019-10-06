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

// ReaderAt is a copy of the standard library io.ReaderAt interface.
type ReaderAt interface {
	io.ReaderAt
}

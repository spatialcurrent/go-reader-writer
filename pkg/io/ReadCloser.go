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

// ReadCloser is a copy of the standard library io.ReadCloser interface.
type ReadCloser interface {
	io.ReadCloser
}

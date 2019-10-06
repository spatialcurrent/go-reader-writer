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

// Closer is a copy of the standard library io.Closer interface.
type Closer interface {
	io.Closer
}

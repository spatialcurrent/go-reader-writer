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

// WriteString writes a string to the writer and returns an error, if any.
func WriteString(w Writer, s string) (int, error) {
	if w == nil {
		return 0, ErrMissingWriter
	}
	return io.WriteString(w, s)
}

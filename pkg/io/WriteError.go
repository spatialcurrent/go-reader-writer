// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package io

// WriteError writes an error with a trailing newline to the writer and returns an error, if any.
func WriteError(w Writer, err error) (int, error) {
	if w == nil {
		return 0, ErrMissingWriter
	}
	return WriteString(w, err.Error()+"\n")
}

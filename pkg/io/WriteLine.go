// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package io

// WriteLine writes a string with a trailing newline to the writer and returns an error, if any.
func WriteLine(w Writer, s string) (int, error) {
	if w == nil {
		return 0, ErrMissingWriter
	}
	return WriteString(w, s+"\n")
}

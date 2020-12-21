// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package bzip2

import (
	"bytes"
	"io"

	"compress/bzip2"
)

// ReadBytes returns a reader for reading bzip2-compressed bytes from an input slice.
func ReadBytes(b []byte) io.Reader {
	// Just uses original reader, since nothing to close.
	return bzip2.NewReader(bytes.NewReader(b))
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package snappy

import (
	"bytes"
	"io"

	"github.com/golang/snappy"
)

// ReadBytes returns a reader for reading snappy-compressed bytes from an input slice.
// b is the input slice of compressed bytes.
//
//  - https://godoc.org/github.com/golang/snappy
//  - https://en.wikipedia.org/wiki/Snappy_(compression)
//
func ReadBytes(b []byte) io.Reader {
	// Just uses original reader, since nothing to close.
	return snappy.NewReader(bytes.NewReader(b))
}

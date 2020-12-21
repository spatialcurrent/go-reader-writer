// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gzip

import (
	"bytes"
	"compress/gzip"
)

// ReadBytes returns a reader for reading zlib-compressed bytes from an input slice.
// b is the input slice of compressed bytes.
// multistream toggles support for a sequence of multiple compressed files with their own header and trailer.
// In most cases, multistream should be true, as it is the standard behavior for a gzip reader.
//
//  - https://golang.org/pkg/compress/zlib/
//  - https://en.wikipedia.org/wiki/Zlib
//
func ReadBytes(b []byte, multistream bool) (ReadResetCloser, error) {
	gr, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	if !multistream {
		gr.Multistream(multistream)
	}
	return gr, nil
}

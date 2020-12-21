// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package flate

import (
	"bytes"
	"compress/flate"
)

// ReadBytes returns a reader for reading DEFLATE-compressed bytes from an input slice.
// b is the input slice of compressed bytes.  dict is the initial dictionary, if one exists.
//
//  - https://golang.org/pkg/compress/flate/
//  - https://en.wikipedia.org/wiki/DEFLATE
//
func ReadBytes(b []byte, dict []byte) ReadResetCloser {
	if len(dict) > 0 {
		return flate.NewReaderDict(bytes.NewReader(b), dict).(ReadResetCloser)
	}

	return flate.NewReader(bytes.NewReader(b)).(ReadResetCloser)
}

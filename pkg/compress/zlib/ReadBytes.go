// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package zlib

import (
	"bytes"
	"compress/zlib"
)

// ReadBytes returns a reader for reading zlib-compressed bytes from an input slice.
// b is the input slice of compressed bytes.  dict is the initial dictionary, if one exists.
//
//  - https://golang.org/pkg/compress/zlib/
//  - https://en.wikipedia.org/wiki/Zlib
//
func ReadBytes(b []byte, dict []byte) (ReadResetCloser, error) {
	if len(dict) > 0 {
		r, err := zlib.NewReaderDict(bytes.NewReader(b), dict)
		if err != nil {
			return nil, err
		}
		return r.(ReadResetCloser), nil
	}

	r, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return r.(ReadResetCloser), nil
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/zlib"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
)

// ReadZlibFile returns a reader for reading bytes from a zlib-compressed file
// Wraps the "compress/zlib" package.
//
//  - https://golang.org/pkg/compress/zlib/
//
func ReadZlibFile(path string, dict []byte, bufferSize int) (*Reader, error) {

	f, err := os.OpenFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening zlib file at path %q for reading", path)
	}

	if len(dict) > 0 {
		zr, err := zlib.NewReaderDict(bufio.NewReaderSize(f, bufferSize), dict)
		if err != nil {
			return nil, errors.Wrapf(err, "error creating zlib reader for file at path %q", path)
		}
		return &Reader{Reader: bufio.NewReader(zr)}, nil
	}

	zr, err := zlib.NewReader(bufio.NewReaderSize(f, bufferSize))
	if err != nil {
		return nil, errors.Wrapf(err, "error creating zlib reader for file at path %q", path)
	}
	return &Reader{Reader: bufio.NewReader(zr)}, nil
}

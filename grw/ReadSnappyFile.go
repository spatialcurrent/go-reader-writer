// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"os"
)

import (
	"github.com/golang/snappy"
	"github.com/pkg/errors"
)

// ReadSnappyFile returns a reader for a snappy-compressed file, and an error if any.
//
//  - https://godoc.org/github.com/golang/snappy
//
func ReadSnappyFile(path string, cache bool, buffer_size int) (ByteReadCloser, error) {

	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening snappy file at \""+path+"\" for reading")
	}

	sr := snappy.NewReader(bufio.NewReaderSize(f, buffer_size))

	if cache {
		return NewCache(&Reader{Reader: bufio.NewReaderSize(sr, buffer_size), File: f}), nil
	}

	return &Reader{Reader: bufio.NewReaderSize(sr, buffer_size), File: f}, nil
}

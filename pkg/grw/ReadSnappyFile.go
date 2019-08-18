// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"io"
)

import (
	"github.com/golang/snappy"
	"github.com/pkg/errors"
)

// ReadSnappyFile returns a reader for a snappy-compressed file, and an error if any.
//
//  - https://godoc.org/github.com/golang/snappy
//
func ReadSnappyFile(path string, buffer_size int) (ByteReadCloser, error) {

	f, err := OpenFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening snappy file")
	}

	return &Reader{Reader: bufio.NewReaderSize(snappy.NewReader(bufio.NewReaderSize(f, buffer_size)), buffer_size), Closer: &Closer{Closers: []io.Closer{f}}}, nil
}

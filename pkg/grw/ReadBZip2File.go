// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"compress/bzip2"
	"io"
)

import (
	"github.com/pkg/errors"
)

// ReadBzip2File returns a reader for reading bytes from a bzip2-compressed file
// Wraps the "compress/gzip" package.
//
//  - https://golang.org/pkg/compress/gzip/
//
func ReadBzip2File(path string, buffer_size int) (ByteReadCloser, error) {

	f, err := OpenFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening bzip2 file")
	}

	return &Reader{Reader: bufio.NewReaderSize(bzip2.NewReader(bufio.NewReaderSize(f, buffer_size)), buffer_size), Closers: []io.Closer{f}}, nil
}

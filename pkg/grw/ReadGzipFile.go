// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"compress/gzip"
	"io"
)

import (
	"github.com/pkg/errors"
)

// ReadGzipFile returns a reader for reading bytes from a gzip-compressed file
// Wraps the "compress/gzip" package.
//
//  - https://golang.org/pkg/compress/gzip/
//
func ReadGzipFile(path string, buffer_size int) (ByteReadCloser, error) {

	f, err := OpenFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening gzip file")
	}

	gr, err := gzip.NewReader(bufio.NewReaderSize(f, buffer_size))
	if err != nil {
		return nil, errors.Wrap(err, "Error creating gzip reader for file \""+path+"\"")
	}

	return &Reader{Reader: bufio.NewReaderSize(gr, buffer_size), Closer: &Closer{Closers: []io.Closer{gr, f}}}, nil
}

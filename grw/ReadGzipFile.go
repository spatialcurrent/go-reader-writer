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
	"os"
)

import (
	"github.com/pkg/errors"
)

// ReadGzipFile returns a reader for reading bytes from a gzip-compressed file
// Wraps the "compress/gzip" package.
//
//  - https://golang.org/pkg/compress/gzip/
//
func ReadGzipFile(path string, cache bool, buffer_size int) (ByteReadCloser, error) {

	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening gzip file at \""+path+"\" for reading")
	}

	gr, err := gzip.NewReader(bufio.NewReaderSize(f, buffer_size))
	if err != nil {
		return nil, errors.Wrap(err, "Error creating gzip reader for file \""+path+"\"")
	}

	if cache {
		return &Cache{
			Reader:  &Reader{Reader: bufio.NewReaderSize(gr, buffer_size), Closer: gr, File: f},
			Content: &[]byte{},
		}, nil
	}

	return &Reader{Reader: bufio.NewReaderSize(gr, buffer_size), Closer: gr, File: f}, nil
}

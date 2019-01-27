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
	"os"
)

import (
	"github.com/pkg/errors"
)

// ReadBzip2File returns a reader for reading bytes from a bzip2-compressed file
// Wraps the "compress/gzip" package.
//
//  - https://golang.org/pkg/compress/gzip/
//
func ReadBzip2File(path string, cache bool, buffer_size int) (ByteReadCloser, error) {

	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening gzip file at \""+path+"\" for reading")
	}

	br := bzip2.NewReader(bufio.NewReaderSize(f, buffer_size))

	if cache {
		return &Cache{
			Reader:  &Reader{Reader: bufio.NewReaderSize(br, buffer_size), File: f},
			Content: &[]byte{},
		}, nil
	}

	return &Reader{Reader: bufio.NewReaderSize(br, buffer_size), File: f}, nil
}

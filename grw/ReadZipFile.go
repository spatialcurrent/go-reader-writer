// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"archive/zip"
	"bufio"
)

import (
	"github.com/pkg/errors"
)

// ReadZipFile returns a ByteReadCloser for reading bytes from a zip-compressed file
// Wraps the "archive/zip" package.
//
//  - https://golang.org/pkg/archive/zip/
//
func ReadZipFile(path string, cache bool) (ByteReadCloser, error) {

	zr, err := zip.OpenReader(path)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening gzip file at \""+path+"\" for reading")
	}

	if len(zr.File) != 1 {
		return nil, errors.New("error zip file has more than one internal file.")
	}

	zfr, err := zr.File[0].Open()
	if err != nil {
		return nil, errors.Wrap(err, "error opening internal file for zip.")
	}

	if cache {
		return &Cache{
			Reader:  &Reader{Reader: bufio.NewReader(zfr), Closer: zfr},
			Content: &[]byte{},
		}, nil
	}

	return &Reader{Reader: bufio.NewReader(zfr), Closer: zfr}, nil
}

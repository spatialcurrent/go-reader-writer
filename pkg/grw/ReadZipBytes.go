// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"archive/zip"
	"bufio"
	"bytes"
	"io"
)

import (
	"github.com/pkg/errors"
)

// ReadZipBytes returns a reader for reading from zip-compressed bytes.
//
//  - https://godoc.org/github.com/golang/snappy
//
func ReadZipBytes(b []byte) (ByteReadCloser, error) {

	zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, errors.Wrap(err, "error creating reader for zip bytes")
	}

	if len(zr.File) != 1 {
		return nil, errors.New("error zip file has more than one internal file")
	}

	zfr, err := zr.File[0].Open()
	if err != nil {
		return nil, errors.Wrap(err, "error opening internal file for zip")
	}

	return &Reader{Reader: bufio.NewReader(zfr), Closers: []io.Closer{zfr}}, nil
}

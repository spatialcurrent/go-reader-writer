// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-reader-writer/pkg/archive/zip"
	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
)

// ReadZipFile returns a ByteReadCloser for reading bytes from a zip-compressed file
// Wraps the "archive/zip" package.
//
//  - https://golang.org/pkg/archive/zip/
//
func ReadZipFile(path string) (*Reader, error) {

	zr, err := zip.OpenReader(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening zip file at path %q for reading", path)
	}

	if len(zr.File) != 1 {
		return nil, errors.New("error zip file has more than one internal file.")
	}

	zfr, err := zr.File[0].Open()
	if err != nil {
		return nil, errors.Wrap(err, "error opening internal file for zip.")
	}

	return &Reader{Reader: bufio.NewReader(zfr)}, nil
}

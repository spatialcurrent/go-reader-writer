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
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/flate"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
)

// ReadFlateFile returns a reader for reading bytes from a DEFLATE-compressed file
// Wraps the "compress/flate" package.
//
//  - https://golang.org/pkg/compress/flate/
//
func ReadFlateFile(path string, dict []byte, bufferSize int) (*Reader, error) {

	f, err := os.OpenFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening flate file at path %q for reading", path)
	}

	if len(dict) > 0 {
		return &Reader{Reader: bufio.NewReader(flate.NewReaderDict(bufio.NewReaderSize(f, bufferSize), dict))}, nil
	}

	return &Reader{Reader: bufio.NewReader(flate.NewReader(bufio.NewReaderSize(f, bufferSize)))}, nil
}

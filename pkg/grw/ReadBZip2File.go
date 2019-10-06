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
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/bzip2"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
)

// ReadBzip2File returns a reader for reading bytes from a bzip2-compressed file
// Wraps the "compress/gzip" package.
//
//  - https://golang.org/pkg/compress/gzip/
//
func ReadBzip2File(path string, bufferSize int) (*Reader, error) {

	f, err := os.OpenFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening bzip2 file")
	}

	return &Reader{Reader: bufio.NewReader(bzip2.NewReader(bufio.NewReaderSize(f, bufferSize)))}, nil
}

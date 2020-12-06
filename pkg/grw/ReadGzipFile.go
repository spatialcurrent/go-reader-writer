// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/gzip"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
)

// ReadGzipFile returns a reader for reading bytes from a gzip-compressed file
// Wraps the "compress/gzip" package.
//
//  - https://golang.org/pkg/compress/gzip/
//
func ReadGzipFile(path string, bufferSize int) (*Reader, error) {

	f, err := os.OpenFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening gzip file at path %q for reading: %w", path, err)
	}

	gr, err := gzip.NewReader(bufio.NewReaderSize(f, bufferSize))
	if err != nil {
		return nil, fmt.Errorf("error creating gzip reader for file at path %q: %w", path, err)
	}

	return &Reader{Reader: bufio.NewReader(gr)}, nil
}

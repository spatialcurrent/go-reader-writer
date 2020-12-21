// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gzip

import (
	"fmt"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
)

// ReadFile returns a reader for reading bytes from a gzip-compressed file.
func ReadFile(path string, bufferSize int) (*Reader, error) {

	f, err := os.OpenFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening gzip file at path %q for reading: %w", path, err)
	}

	gr, err := NewReader(bufio.NewReaderSize(f, bufferSize))
	if err != nil {
		return nil, fmt.Errorf("error creating gzip reader for file at path %q: %w", path, err)
	}

	return gr, nil
}

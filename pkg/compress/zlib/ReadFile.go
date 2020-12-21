// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package zlib

import (
	"fmt"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
)

// ReadFile returns a reader for reading bytes from a zlib-compressed file.
func ReadFile(path string, dict []byte, bufferSize int) (*Reader, error) {

	f, err := os.OpenFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening zlib file at path %q for reading: %w", path, err)
	}

	if len(dict) > 0 {
		zr, err := NewReaderDict(bufio.NewReaderSize(f, bufferSize), dict)
		if err != nil {
			return nil, fmt.Errorf("error creating zlib reader for file at path %q: %w", path, err)
		}
		return zr, nil
	}

	zr, err := NewReader(bufio.NewReaderSize(f, bufferSize))
	if err != nil {
		return nil, fmt.Errorf("error creating zlib reader for file at path %q: %w", path, err)
	}

	return zr, nil
}

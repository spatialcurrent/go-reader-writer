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

	f, errOpenFile := os.OpenFile(path)
	if errOpenFile != nil {
		return nil, fmt.Errorf("error opening zlib file at path %q for reading: %w", path, errOpenFile)
	}

	if len(dict) > 0 {
		zr, errNewReader := NewReaderDict(bufio.NewReaderSize(f, bufferSize), dict)
		if errNewReader != nil {
			return nil, fmt.Errorf("error creating zlib reader for file at path %q: %w", path, errNewReader)
		}
		return zr, nil
	}

	zr, errNewReader := NewReader(bufio.NewReaderSize(f, bufferSize))
	if errNewReader != nil {
		return nil, fmt.Errorf("error creating zlib reader for file at path %q: %w", path, errNewReader)
	}

	return zr, nil
}

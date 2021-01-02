// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package zlib

import (
	"fmt"
	"os"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
)

// WriteFile returns a Writer for writing to a local file
func WriteFile(path string, dict []byte, bufferSize int) (*Writer, error) {
	if bufferSize < 0 {
		return nil, fmt.Errorf("error creating zlib writer for file at path %q: invalid buffer size %d", path, bufferSize)
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("error opening file at path %q for writing: %w", path, err)
	}
	if bufferSize > 0 {
		w, errWriter := NewWriterDict(bufio.NewWriterSize(f, bufferSize), dict)
		if errWriter != nil {
			return nil, fmt.Errorf("error creating zlib writer for file at path %q: %w", path, errWriter)
		}
		return w, nil
	}
	w, errWriter := NewWriterDict(f, dict)
	if errWriter != nil {
		return nil, fmt.Errorf("error creating zlib writer for file at path %q: %w", path, errWriter)
	}
	return w, nil
}

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
	"os"
)

// WriteFile returns a Writer for writing to a local file
func WriteFile(path string, bufferSize int) (*Writer, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("error opening file at path %q for writing: %w", path, err)
	}
	if bufferSize > 0 {
		return NewWriter(bufio.NewWriterSize(f, bufferSize)), nil
	}
	if bufferSize < 0 {
		return nil, fmt.Errorf("error creating GZIP writer for file at path %q: invalid buffer size %d", path, bufferSize)
	}
	return NewWriter(f), nil
}

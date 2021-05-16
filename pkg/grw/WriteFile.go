// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
)

// WriteFileInput holds the input for the WriteFile method
type WriteFileInput struct {
	Path       string // required field
	BufferSize int    // if zero, then no buffer is used.
	Mode       uint32 // defaults to 0600
}

// WriteFile returns a Writer for writing to a local file
func WriteFile(input *WriteFileInput) (io.WriteCloser, error) {
	if input == nil {
		return nil, errors.New("input is nil")
	}
	path := input.Path
	if len(path) == 0 {
		return nil, errors.New("invalid input: path is nil")
	}
	bufferSize := input.BufferSize
	mode := input.Mode
	if mode == uint32(0) {
		mode = uint32(0600)
	}
	if bufferSize < 0 {
		return nil, fmt.Errorf("error creating writer for file at path %q: invalid buffer size %d", path, bufferSize)
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.FileMode(mode))
	if err != nil {
		return nil, fmt.Errorf("error opening file at path %q for writing: %w", path, err)
	}
	if bufferSize > 0 {
		return bufio.NewWriterSize(f, bufferSize), nil
	}
	return f, nil
}

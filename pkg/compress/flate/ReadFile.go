// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package flate

import (
	"fmt"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
)

// ReadFile returns a reader for reading bytes from a DEFLATE-compressed file.
func ReadFile(path string, dict []byte, bufferSize int) (*Reader, error) {

	f, err := os.OpenFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening flate file at path %q for reading: %w", path, err)
	}

	if len(dict) > 0 {
		return NewReaderDict(bufio.NewReaderSize(f, bufferSize), dict), nil
	}

	return NewReader(bufio.NewReaderSize(f, bufferSize)), nil
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
	"os"
)

// WriteFlateFile returns a ByteWriteCloser for writing to a local file
func WriteFlateFile(path string, dict []byte, flag int) (*Writer, error) {
	f, err := os.OpenFile(path, flag, 0600)
	if err != nil {
		return nil, fmt.Errorf("error opening file at path %q for writing: %w", path, err)
	}
	w, err := WrapWriter(f, AlgorithmFlate, dict)
	if err != nil {
		return nil, fmt.Errorf("error wrapping writer: %w", err)
	}
	return w, nil
}

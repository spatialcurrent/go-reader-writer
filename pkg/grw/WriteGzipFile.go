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

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/gzip"
)

// WriteGzipFile returns a ByteWriteCloser for writing to a local file
func WriteGzipFile(path string, flag int) (*Writer, error) {
	f, err := os.OpenFile(path, flag, 0600)
	if err != nil {
		return nil, fmt.Errorf("error opening file at path %q for writing: %w", path, err)
	}
	return NewWriter(bufio.NewWriter(gzip.NewWriter(f))), nil
}

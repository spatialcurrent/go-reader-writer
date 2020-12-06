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
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/snappy"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
)

// ReadSnappyFile returns a reader for a snappy-compressed file, and an error if any.
//
//  - https://godoc.org/github.com/golang/snappy
//
func ReadSnappyFile(path string, bufferSize int) (*Reader, error) {

	f, err := os.OpenFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening snappy file: %w", err)
	}

	return &Reader{Reader: bufio.NewReader(snappy.NewReader(bufio.NewReaderSize(f, bufferSize)))}, nil
}

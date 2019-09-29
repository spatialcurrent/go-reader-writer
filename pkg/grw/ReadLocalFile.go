// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
)

// ReadLocalFile returns a ByteReader for reading bytes without any transformation from a file, and an error if any.
func ReadLocalFile(path string, bufferSize int) (*Reader, error) {

	f, err := os.OpenFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening regular file")
	}

	return &Reader{Reader: bufio.NewReaderSize(f, bufferSize)}, nil
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"os"

	"github.com/golang/snappy"
	"github.com/pkg/errors"
)

// WriteSnappyFile returns a ByteWriteCloser for writing to a local file
func WriteSnappyFile(path string, flag int) (ByteWriteCloser, error) {

	f, err := os.OpenFile(path, flag, 0600)
	if err != nil {
		return nil, errors.Wrap(err, "error opening file at \""+path+"\" for writing")
	}

	sw := snappy.NewBufferedWriter(f)

	return NewWriterWithCloserAndFile(bufio.NewWriter(sw), sw, f), nil
}

// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"os"
)

import (
	"github.com/pkg/errors"
)

// WriteLocalFile returns a ByteWriteCloser for writing to a local file
func WriteLocalFile(path string, flag int) (ByteWriteCloser, error) {

	f, err := os.OpenFile(path, flag, 0600)
	if err != nil {
		return nil, errors.Wrap(err, "error opening file at \""+path+"\" for writing")
	}

	return &Writer{Writer: bufio.NewWriter(f), File: f}, nil
}

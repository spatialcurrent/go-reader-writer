// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// WriteLocalFile returns a ByteWriteCloser for writing to a local file
func WriteLocalFile(path string, flag int, parents bool) (ByteWriteCloser, error) {

	if parents {
		err := os.MkdirAll(filepath.Dir(path), 0700)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error creating parent directories for %q", path))
		}
	}

	f, err := os.OpenFile(path, flag, 0600)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error opening file for writing at path %q", path))
	}

	return NewWriterWithCloserAndFile(bufio.NewWriter(f), nil, f), nil
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"os"

	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
)

// WriteLocalFile returns a ByteWriteCloser for writing to a local file
func WriteLocalFile(path string, flag int) (*Writer, error) {

	f, err := os.OpenFile(path, flag, 0600)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening file for writing at path %q", path)
	}

	return NewWriter(bufio.NewWriter(f)), nil
}

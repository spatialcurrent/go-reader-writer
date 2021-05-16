// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package zip

import (
	"errors"
	"fmt"
	"io"
	//"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
)

// ReadFile returns a Reader for reading bytes from a zip-compressed file.
func ReadFile(path string) (io.ReadCloser, error) {

	zr, err := OpenReader(path)
	if err != nil {
		return nil, fmt.Errorf("error opening zip file at path %q for reading: %w", path, err)
	}

	if len(zr.File) != 1 {
		return nil, errors.New("error zip file has more than one internal file")
	}

	zfr, err := zr.File[0].Open()
	if err != nil {
		return nil, fmt.Errorf("error opening internal file for zip: %w", err)
	}

	return zfr, nil
}

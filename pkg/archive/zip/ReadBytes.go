// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package zip

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
)

// ReadBytes returns a reader for reading zip-compressed bytes from an input slice.
// b is the input slice of compressed bytes.
//
//  - https://golang.org/pkg/archive/zip
//
func ReadBytes(b []byte) (io.ReadCloser, error) {

	zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, fmt.Errorf("error creating reader for zip bytes: %w", err)
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

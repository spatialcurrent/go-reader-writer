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

	"github.com/spatialcurrent/go-reader-writer/pkg/archive/zip"
)

// ReadZipBytes returns a reader for reading from zip-compressed bytes.
//
//  - https://godoc.org/github.com/golang/snappy
//
func ReadZipBytes(b []byte) (*Reader, error) {

	zfr, err := zip.ReadBytes(b)
	if err != nil {
		return nil, fmt.Errorf("error reading zip bytes: %w", err)
	}

	return &Reader{Reader: bufio.NewReader(zfr)}, nil
}

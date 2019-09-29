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
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/gzip"
)

// GzipBytes returns a reader for reading gzip bytes from an input array.
// Wraps the "compress/gzip" package.
//
//  - https://golang.org/pkg/compress/gzip/
//
func ReadGzipBytes(b []byte) (*Reader, error) {
	gr, err := gzip.ReadBytes(b, true)
	if err != nil {
		return nil, errors.Wrap(err, "error creating gzip reader for memory block.")
	}
	return &Reader{Reader: bufio.NewReader(gr)}, nil
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"io"
)

import (
	"github.com/pkg/errors"
)

// GzipBytes returns a reader for reading gzip bytes from an input array.
// Wraps the "compress/gzip" package.
//
//  - https://golang.org/pkg/compress/gzip/
//
func ReadGzipBytes(b []byte) (ByteReadCloser, error) {
	gr, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "error creating gzip reader for memory block.")
	}
	return &Reader{Reader: bufio.NewReader(gr), Closer: &Closer{Closers: []io.Closer{gr}}}, nil
}

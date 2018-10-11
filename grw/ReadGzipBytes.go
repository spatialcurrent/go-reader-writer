// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"bytes"
	"compress/gzip"
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

	r := &Cache{
		Reader: &Reader{
			Reader: bufio.NewReader(gr),
			Closer: gr,
		},
		Content: &[]byte{},
	}

	return r, nil
}

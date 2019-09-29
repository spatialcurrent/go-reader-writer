// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package bzip2

import (
	"io"

	"compress/bzip2"

	"github.com/pkg/errors"
)

type Reader struct {
	io.Reader
	underlying io.Reader
}

// Close closes the underlying reader if it implements io.Closer.
func (r *Reader) Close() error {
	if c, ok := r.underlying.(io.Closer); ok {
		err := c.Close()
		if err != nil {
			return errors.Wrap(err, "error closing underlying reader")
		}
	}
	return nil
}

// NewReader returns an io.Reader which decompresses bzip2 data from r.
// If r does not also implement io.ByteReader, the decompressor may read more data than necessary from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{Reader: bzip2.NewReader(r), underlying: r}
}

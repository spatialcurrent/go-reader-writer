// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package bufio

import (
	"bufio"
	"io"

	"github.com/pkg/errors"
)

type Reader struct {
	*bufio.Reader
	underlying io.Reader
}

// Close, calls the Close method of the underlying reader, if it implements io.Closer.
func (r *Reader) Close() error {
	if c, ok := r.underlying.(io.Closer); ok {
		err := c.Close()
		if err != nil {
			return errors.Wrap(err, "error closing underlying reader")
		}
	}
	return nil
}

// NewReader returns a new Reader whose buffer has the default size.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		Reader:     bufio.NewReader(r),
		underlying: r,
	}
}

// NewReaderSize returns a new Reader whose buffer has at least the specified size. If the argument io.Reader is already a Reader with large enough size, it returns the underlying Reader.
func NewReaderSize(r io.Reader, size int) *Reader {
	return &Reader{
		Reader:     bufio.NewReaderSize(r, size),
		underlying: r,
	}
}

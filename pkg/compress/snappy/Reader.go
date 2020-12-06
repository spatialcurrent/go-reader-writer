// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package snappy

import (
	"fmt"
	"io"

	"github.com/golang/snappy"
)

type Reader struct {
	*snappy.Reader
	underlying io.Reader
}

// Close closes the underlying reader if it implements io.Closer.
func (r *Reader) Close() error {
	if c, ok := r.underlying.(io.Closer); ok {
		err := c.Close()
		if err != nil {
			return fmt.Errorf("error closing underlying reader: %w", err)
		}
	}
	return nil
}

// Reset discards any buffered data, resets all state, and switches the Snappy reader to read from r.
// This permits reusing a Reader rather than allocating a new one.
func (r *Reader) Reset(reader io.Reader) {
	r.Reader.Reset(reader)
	r.underlying = reader
}

// NewReader returns a new Reader that decompresses from r, using the framing format described at https://github.com/google/snappy/blob/master/framing_format.txt
func NewReader(r io.Reader) *Reader {
	return &Reader{Reader: snappy.NewReader(r), underlying: r}
}

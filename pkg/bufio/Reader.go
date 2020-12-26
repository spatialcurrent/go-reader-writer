// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package bufio

import (
	"bufio"
	"fmt"
	"io"
)

// A reader that propagates the close method.
type Reader struct {
	*bufio.Reader
	underlying io.Closer
	close      bool
}

func (b *Reader) Reset(r io.ReadCloser) {
	b.Reader.Reset(r)
	b.underlying = r
}

func (b *Reader) Close() error {
	if b.close {
		err := b.underlying.Close()
		if err != nil {
			return fmt.Errorf("error closing underlying reader: %w", err)
		}
	}
	return nil
}

// NewReader returns a new NewReader whose buffer has the default size.
func NewReader(r io.ReadCloser) *Reader {
	return &Reader{
		Reader:     bufio.NewReader(r),
		underlying: r,
		close:      true,
	}
}

// NewReaderSize returns a new Reader whose buffer has at least the specified size. If the argument io.Reader is already a Reader with large enough size, it returns the underlying Reader.
func NewReaderSize(r io.ReadCloser, size int) *Reader {
	return &Reader{
		Reader:     bufio.NewReaderSize(r, size),
		underlying: r,
		close:      true,
	}
}

// NewReaderSize returns a new Reader whose buffer has at least the specified size. If the argument io.Reader is already a Reader with large enough size, it returns the underlying Reader.
func NewReaderSizeClose(r io.ReadCloser, size int, close bool) *Reader {
	return &Reader{
		Reader:     bufio.NewReaderSize(r, size),
		underlying: r,
		close:      close,
	}
}

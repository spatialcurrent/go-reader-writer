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

type Writer struct {
	*bufio.Writer
	underlying io.Writer
	close      bool
}

type flusher interface {
	Flush() error
}

// Flush writes any buffered data to the underlying io.Writer.  Then calls the "Flush() error" method of the underlying writer, if it implements it.
func (b *Writer) Flush() error {
	err := b.Writer.Flush()
	if err != nil {
		return fmt.Errorf("error flushing bufio.Writer: %w", err)
	}
	if f, ok := b.underlying.(flusher); ok {
		err = f.Flush()
		if err != nil {
			return fmt.Errorf("error flushing underlying writer: %w", err)
		}
	}
	return nil
}

// Close, calls the "Close() error" method of the underlying writer, if it implements io.Closer.
func (b *Writer) Close() error {
	if b.close {
		if c, ok := b.underlying.(io.Closer); ok {
			err := c.Close()
			if err != nil {
				return fmt.Errorf("error closing underlying writer: %w", err)
			}
		}
	}
	return nil
}

// NewWriter returns a new Writer whose buffer has the default size.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Writer:     bufio.NewWriter(w),
		underlying: w,
		close:      true,
	}
}

// NewWriter returns a new Writer whose buffer has the default size.  If close is false, then does not close the underlying writer.
func NewWriterClose(w io.Writer, close bool) *Writer {
	return &Writer{
		Writer:     bufio.NewWriter(w),
		underlying: w,
		close:      close,
	}
}

// NewWriterSize returns a new Writer whose buffer has at least the specified size. If the argument io.Writer is already a Writer with large enough size, it returns the underlying Writer.
func NewWriterSize(w io.Writer, size int) *Writer {
	return &Writer{
		Writer:     bufio.NewWriterSize(w, size),
		underlying: w,
		close:      true,
	}
}

// NewWriterSize returns a new Writer whose buffer has at least the specified size. If the argument io.Writer is already a Writer with large enough size, it returns the underlying Writer.
func NewWriterSizeClose(w io.Writer, size int, close bool) *Writer {
	return &Writer{
		Writer:     bufio.NewWriterSize(w, size),
		underlying: w,
		close:      close,
	}
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package snappy

import (
	"io"

	"github.com/golang/snappy"

	"github.com/pkg/errors"
)

type Writer struct {
	*snappy.Writer
	underlying io.Writer
}

type flusher interface {
	Flush() error
}

// Close, flushes and closes the snppy.Writer and calls the "Close() error" method of the underlying writer, if it implements io.Closer.
func (w *Writer) Close() error {
	err := w.Writer.Close()
	if err != nil {
		return errors.Wrap(err, "error closing snappy writer")
	}
	// When the snappy writer is closed is flushes one last time.
	// Therefore, we need to flush the underlying writer one last time before we close it.
	if f, ok := w.underlying.(flusher); ok {
		err = f.Flush()
		if err != nil {
			return errors.Wrap(err, "error flushing underlying writer")
		}
	}
	if c, ok := w.underlying.(io.Closer); ok {
		err = c.Close()
		if err != nil {
			return errors.Wrap(err, "error closing underlying writer")
		}
	}
	return nil
}

// Flush writes any pending data to the underlying io.Writer.  Then calls the "Flush() error" method of the underlying writer, if it implements it.
func (w *Writer) Flush() error {
	err := w.Writer.Flush()
	if err != nil {
		return errors.Wrap(err, "error flushing snappy writer")
	}
	if f, ok := w.underlying.(flusher); ok {
		err = f.Flush()
		if err != nil {
			return errors.Wrap(err, "error flushing underlying writer")
		}
	}
	return nil
}

// Reset discards the writer's state and switches the Snappy writer to write to w. This permits reusing a Writer rather than allocating a new one
func (w *Writer) Reset(writer io.Writer) {
	w.Writer.Reset(writer)
	w.underlying = w
}

// NewBufferedWriter returns a new Writer that compresses to w, using the
// framing format described at
// https://github.com/google/snappy/blob/master/framing_format.txt
//
// The Writer returned buffers writes. Users must call Close to guarantee all
// data has been forwarded to the underlying io.Writer. They may also call
// Flush zero or more times before calling Close.
func NewBufferedWriter(w io.Writer) *Writer {
	return &Writer{Writer: snappy.NewBufferedWriter(w), underlying: w}
}

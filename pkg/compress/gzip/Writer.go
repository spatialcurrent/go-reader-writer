// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package gzip provides a writer that embeds the standard library gzip.Writer and propagates calls to Flush and Close.
package gzip

import (
	"compress/gzip"
	"fmt"
	"io"
)

type Writer struct {
	*gzip.Writer
	underlying io.Writer
}

type flusher interface {
	Flush() error
}

// Flush flushes any pending compressed data to the underlying writer.
//
// It is useful mainly in compressed network protocols, to ensure that
// a remote reader has enough data to reconstruct a packet. Flush does
// not return until the data has been written. If the underlying
// writer returns an error, Flush returns that error.
//
// In the terminology of the zlib library, Flush is equivalent to Z_SYNC_FLUSH.
func (w *Writer) Flush() error {
	err := w.Writer.Flush()
	if err != nil {
		return fmt.Errorf("error flushing gzip writer: %w", err)
	}
	if f, ok := w.underlying.(flusher); ok {
		err = f.Flush()
		if err != nil {
			return fmt.Errorf("error flushing underlying writer: %w", err)
		}
	}
	return nil
}

// Close closes the Writer by flushing any unwritten data to the underlying io.Writer and writing the GZIP footer.
// Calls the "Close() error" method of the underlying writer, if it implements io.Closer.
func (w *Writer) Close() error {
	err := w.Writer.Close()
	if err != nil {
		return fmt.Errorf("error closing gzip writer: %w", err)
	}
	// When the gzip writer is closed is writes one final trailer to the underlying writer.
	// Therefore, we need to flush the underlying writer one last time before we close it.
	if f, ok := w.underlying.(flusher); ok {
		err = f.Flush()
		if err != nil {
			return fmt.Errorf("error flushing underlying writer: %w", err)
		}
	}
	if c, ok := w.underlying.(io.Closer); ok {
		err = c.Close()
		if err != nil {
			return fmt.Errorf("error closing underlying writer: %w", err)
		}
	}
	return nil
}

// Reset discards the Writer z's state and makes it equivalent to the
// result of its original state from NewWriter or NewWriterLevel, but
// writing to w instead. This permits reusing a Writer rather than
// allocating a new one.
func (w *Writer) Reset(writer io.Writer) {
	w.Writer.Reset(writer)
	w.underlying = writer
}

// NewWriter returns a new Writer.
// Writes to the returned writer are compressed and written to w.
//
// It is the caller's responsibility to call Close on the Writer when done.
// Writes may be buffered and not flushed until Close.
//
// Callers that wish to set the fields in Writer.Header must do so before
// the first call to Write, Flush, or Close.
func NewWriter(w io.Writer) *Writer {
	return &Writer{Writer: gzip.NewWriter(w), underlying: w}
}

// NewWriterLevel is like NewWriter but specifies the compression level instead
// of assuming DefaultCompression.
//
// The compression level can be DefaultCompression, NoCompression, HuffmanOnly
// or any integer value between BestSpeed and BestCompression inclusive.
// The error returned will be nil if the level is valid.
func NewWriterLevel(w io.Writer, level int) (*Writer, error) {
	gw, err := gzip.NewWriterLevel(w, level)
	if err != nil {
		return nil, err
	}
	return &Writer{Writer: gw, underlying: w}, nil
}

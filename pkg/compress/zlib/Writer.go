// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package zlib provides a writer that embeds the standard library zlib.Writer and propagates calls to Flush and Close.
package zlib

import (
	"compress/zlib"
	"fmt"
	"io"
)

const (
	NoCompression      = zlib.NoCompression
	BestSpeed          = zlib.BestSpeed
	BestCompression    = zlib.BestCompression
	DefaultCompression = zlib.DefaultCompression
	HuffmanOnly        = zlib.HuffmanOnly
)

type Writer struct {
	*zlib.Writer
	underlying io.WriteCloser
}

type flusher interface {
	Flush() error
}

// Close closes the zlib writer, and then flushes and closes the underlying writer.
func (w *Writer) Close() error {
	err := w.Writer.Close()
	if err != nil {
		return fmt.Errorf("error closing zlib writer: %w", err)
	}
	// When the zlib writer is closed is flushes one last time.
	// Therefore, we need to flush the underlying writer one last time before we close it.
	if f, ok := w.underlying.(flusher); ok {
		err = f.Flush()
		if err != nil {
			return fmt.Errorf("error flushing underlying writer: %w", err)
		}
	}
	err = w.underlying.Close()
	if err != nil {
		return fmt.Errorf("error closing underlying writer: %w", err)
	}
	return nil
}

// Flush writes any pending data to the underlying writer.  Then calls the Flush method of the underlying writer, if it implements it.
func (w *Writer) Flush() error {
	err := w.Writer.Flush()
	if err != nil {
		return fmt.Errorf("error flushing zlib writer: %w", err)
	}
	if f, ok := w.underlying.(flusher); ok {
		err = f.Flush()
		if err != nil {
			return fmt.Errorf("error flushing underlying writer: %w", err)
		}
	}
	return nil
}

// NewWriter creates a new Writer.
// Writes to the returned Writer are compressed and written to w.
//
// It is the caller's responsibility to call Close on the Writer when done.
// Writes may be buffered and not flushed until Close.
func NewWriter(w io.WriteCloser) *Writer {
	return &Writer{Writer: zlib.NewWriter(w), underlying: w}
}

// NewWriterLevel is like NewWriter but specifies the compression level instead
// of assuming DefaultCompression.
//
// The compression level can be DefaultCompression, NoCompression, HuffmanOnly
// or any integer value between BestSpeed and BestCompression inclusive.
// The error returned will be nil if the level is valid.
func NewWriterLevel(w io.WriteCloser, level int) (*Writer, error) {
	zw, err := zlib.NewWriterLevel(w, level)
	if err != nil {
		return nil, err
	}
	return &Writer{Writer: zw, underlying: w}, nil
}

// NewWriterDict is like NewWriter but specifies a dictionary to
// compress with.
//
// The dictionary may be nil. If not, its contents should not be modified until
// the Writer is closed.
func NewWriterDict(w io.WriteCloser, dict []byte) (*Writer, error) {
	zw, err := zlib.NewWriterLevelDict(w, zlib.DefaultCompression, dict)
	if err != nil {
		return nil, err
	}
	return &Writer{Writer: zw, underlying: w}, nil
}

// NewWriterLevelDict is like NewWriterLevel but specifies a dictionary to
// compress with.
//
// The dictionary may be nil. If not, its contents should not be modified until
// the Writer is closed.
func NewWriterLevelDict(w io.WriteCloser, level int, dict []byte) (*Writer, error) {
	zw, err := zlib.NewWriterLevelDict(w, level, dict)
	if err != nil {
		return nil, err
	}
	return &Writer{Writer: zw, underlying: w}, nil
}

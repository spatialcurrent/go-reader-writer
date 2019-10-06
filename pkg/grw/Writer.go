// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"sync"

	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-reader-writer/pkg/io"
)

// Writer is a struct for normalizing reading of bytes from files with arbitrary compression and for closing underlying resources.
// Writer implements the ByteWriter interface by wrapping around a subordinate ByteWriter.
type Writer struct {
	*sync.Mutex               // inherits Lock and Unlock Functions
	Writer      io.ByteWriter // the instance of ByteWriter used for reading bytes
}

func NewWriter(w io.ByteWriter) *Writer {
	return &Writer{Writer: w, Mutex: &sync.Mutex{}}
}

// WriteString writes a slice of bytes to the underlying writer and returns an error, if any.
//  - https://godoc.org/io#Writer
func (w *Writer) Write(p []byte) (n int, err error) {

	if w.Writer != nil {
		return w.Writer.Write(p)
	}

	return 0, nil
}

// WriteByte writes a single byte to the underlying writer.
func (w *Writer) WriteByte(b byte) error {

	if w.Writer != nil {
		return w.Writer.WriteByte(b)
	}

	return nil
}

// WriteString writes a string to the underlying writer and returns an error, if any.
//  - https://godoc.org/io#Writer
func (w *Writer) WriteString(s string) (n int, err error) {

	if w.Writer != nil {
		return io.WriteString(w.Writer, s)
	}

	return 0, nil
}

// WriteLine writes a string with a trailing newline to the underlying writer and returns an error, if any.
//  - https://godoc.org/io#Writer
func (w *Writer) WriteLine(s string) (n int, err error) {

	if w.Writer != nil {
		return io.WriteLine(w.Writer, s)
	}

	return 0, nil
}

// WriteLineSafe writes a string with a trailing newline to the underlying writer and returns an error, if any.
// WriteLineSafe also locks the writer for the duration of writing using a sync.Mutex.
//  - https://godoc.org/io#Writer
//  - https://godoc.org/sync#Mutex
func (w *Writer) WriteLineSafe(s string) (n int, err error) {

	if w.Writer != nil {
		w.Lock()
		n, err := io.WriteLine(w.Writer, s)
		w.Unlock()
		return n, err
	}

	return 0, nil
}

// WriteError writes a an error as a string with a trailing newline to the underlying writer and returns an error, if any.
//  - https://godoc.org/io#Writer
func (w *Writer) WriteError(e error) (n int, err error) {

	if w.Writer != nil {
		return io.WriteError(w.Writer, e)
	}

	return 0, nil
}

// WriteError writes a an error as a string with a trailing newline to the underlying writer and returns an error, if any.
// WriteErrorSafe also locks the writer for the duration of writing using a sync.Mutex.
//  - https://godoc.org/io#Writer
//  - https://godoc.org/sync#Mutex
func (w *Writer) WriteErrorSafe(e error) (n int, err error) {

	if w.Writer != nil {
		w.Lock()
		n, err := io.WriteError(w.Writer, e)
		w.Unlock()
		return n, err
	}

	return 0, nil
}

// Flush recursively flushes all the underlying writers.
//  - https://godoc.org/io#Writer
func (w *Writer) Flush() error {
	if f, ok := w.Writer.(io.Flusher); ok {
		err := f.Flush()
		if err != nil {
			return errors.Wrapf(err, "error flushing underlying writer")
		}
	}
	return nil
}

// FlushSafe flushes any intermediate writer.
// FlushSafe also locks the writer for the duration of flushing using a sync.Mutex.
//  - https://godoc.org/io#Writer
//  - https://godoc.org/sync#Mutex
func (w *Writer) FlushSafe() error {

	if f, ok := w.Writer.(io.Flusher); ok {
		w.Lock()
		err := f.Flush()
		if err != nil {
			w.Unlock()
			return errors.Wrapf(err, "error flushing underlying writer")
		}
		w.Unlock()
	}

	return nil
}

// Close recursively closes all the underlying writers.
func (w *Writer) Close() error {
	if c, ok := w.Writer.(io.Closer); ok {
		err := c.Close()
		if err != nil {
			return errors.Wrapf(err, "error closing underlying writer")
		}
	}
	return nil
}

// CloseSafe closes the Closer and the underlying *os.File if not nil.
// CloseSafe also locks the writer for the duration of flushing using a sync.Mutex.
//  - https://godoc.org/os#File
//  - https://godoc.org/sync#Mutex
func (w *Writer) CloseSafe() error {
	w.Lock()
	err := w.Close()
	w.Unlock()
	return err
}

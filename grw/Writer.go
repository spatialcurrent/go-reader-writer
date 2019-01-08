// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"io"
	"os"
	"sync"
)

import (
	"github.com/pkg/errors"
)

// Writer is a struct for normalizing reading of bytes from files with arbitrary compression and for closing underlying resources.
// Writer implements the ByteWriter interface by wrapping around a subordinate ByteWriter.
type Writer struct {
	*sync.Mutex            // inherits Lock and Unlock Functions
	Writer      ByteWriter // the instance of ByteWriter used for reading bytes
	Closer      io.Closer  // Used for closing readers with footer metadata, e.g., gzip.  Not always needed, e.g., snappy
	File        *os.File   // underlying file, if any
}

func NewWriter(writer ByteWriter) *Writer {
	return &Writer{
		Writer: writer,
		Closer: nil,
		File:   nil,
		Mutex:  &sync.Mutex{}}
}

func NewWriterWithCloser(writer ByteWriter, closer io.Closer) *Writer {
	return &Writer{
		Writer: writer,
		Closer: closer,
		File:   nil,
		Mutex:  &sync.Mutex{}}
}

func NewWriterWithCloserAndFile(writer ByteWriter, closer io.Closer, file *os.File) *Writer {
	return &Writer{
		Writer: writer,
		Closer: closer,
		File:   nil,
		Mutex:  &sync.Mutex{}}
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
		return io.WriteString(w.Writer, s+"\n")
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
		n, err := io.WriteString(w.Writer, s+"\n")
		w.Unlock()
		return n, err
	}

	return 0, nil
}

// WriteError writes a an error as a string with a trailing newline to the underlying writer and returns an error, if any.
//  - https://godoc.org/io#Writer
func (w *Writer) WriteError(e error) (n int, err error) {

	if w.Writer != nil {
		return io.WriteString(w.Writer, e.Error()+"\n")
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
		n, err := io.WriteString(w.Writer, e.Error()+"\n")
		w.Unlock()
		return n, err
	}

	return 0, nil
}

// Flush flushes any intermediate writer.
//  - https://godoc.org/io#Writer
func (w *Writer) Flush() error {

	if w.Writer != nil {
		err := w.Writer.Flush()
		if err != nil {
			return errors.Wrap(err, "error flushing underlying writer")
		}
	}

	return nil
}

// FlushSafe flushes any intermediate writer.
// FlushSafe also locks the writer for the duration of flushing using a sync.Mutex.
//  - https://godoc.org/io#Writer
//  - https://godoc.org/sync#Mutex
func (w *Writer) FlushSafe() error {

	if w.Writer != nil {
		w.Lock()
		err := w.Writer.Flush()
		w.Unlock()
		if err != nil {
			return errors.Wrap(err, "error flushing underlying writer")
		}
	}

	return nil
}

// Close closes the Closer and the underlying *os.File if not nil.
//  - https://godoc.org/os#File
func (w *Writer) Close() error {

	err := w.Flush()
	if err != nil {
		return errors.Wrap(err, "error flushing writer")
	}

	if w.Closer != nil {
		err := w.Closer.Close()
		if err != nil {
			return errors.Wrap(err, "error closing write closer.")
		}
	}

	return w.CloseFile()
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

// CloseFile closes the underlying file and bypasses the writer to stop writing immediately.
func (w *Writer) CloseFile() error {

	if w.File != nil {
		err := w.File.Close()
		if err != nil {
			return errors.Wrap(err, "error closing underlying file.")
		}
	}

	return nil
}

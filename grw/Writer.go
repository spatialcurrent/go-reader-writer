// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"io"
	//"io/ioutil"
	"os"
)

import (
	"github.com/pkg/errors"
)

// Writer is a struct for normalizing reading of bytes from files with arbitrary compression and for closing underlying resources.
// Writer implements the ByteWriter interface by wrapping around a subordinate ByteWriter.
type Writer struct {
	Writer ByteWriter // the instance of ByteWriter used for reading bytes
	Closer io.Closer  // Used for closing readers with footer metadata, e.g., gzip.  Not always needed, e.g., snappy
	File   *os.File   // underlying file, if any
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

// WriteError writes a an error as a string with a trailing newline to the underlying writer and returns an error, if any.
//  - https://godoc.org/io#Writer
func (w *Writer) WriteError(e error) (n int, err error) {

	if w.Writer != nil {
		return io.WriteString(w.Writer, e.Error()+"\n")
	}

	return 0, nil
}

// Flush flushes any intermediate writer
func (w *Writer) Flush() error {

	if w.Writer != nil {
		err := w.Writer.Flush()
		if err != nil {
			return errors.Wrap(err, "error flushing underlying writer")
		}
	}

	return nil
}

// Close closes the Closer and the underlying *os.File if not nil.
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

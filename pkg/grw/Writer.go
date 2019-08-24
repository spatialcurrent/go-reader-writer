// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"io"
	"sync"

	"github.com/pkg/errors"
)

// Writer is a struct for normalizing reading of bytes from files with arbitrary compression and for closing underlying resources.
// Writer implements the ByteWriter interface by wrapping around a subordinate ByteWriter.
type Writer struct {
	*sync.Mutex            // inherits Lock and Unlock Functions
	Writer      ByteWriter // the instance of ByteWriter used for reading bytes
	Flushers    []Flusher
	Closer      *Closer // the underlying closers
}

func NewBufferedWriterWithClosers(writer io.Writer, closers ...io.Closer) *Writer {
	bw := bufio.NewWriter(writer)
	flushers := []Flusher{bw}
	if f, ok := writer.(Flusher); ok {
		flushers = append(flushers, f)
	}
	return &Writer{
		Writer:   bw,
		Flushers: flushers,
		Closer: &Closer{
			Closers: closers,
		},
		Mutex: &sync.Mutex{},
	}
}

func NewBufferedWriter(writer io.Writer) *Writer {
	bw := bufio.NewWriter(writer)
	flushers := []Flusher{bw}
	if f, ok := writer.(Flusher); ok {
		flushers = append(flushers, f)
	}
	return &Writer{
		Writer:   bw,
		Flushers: flushers,
		Closer:   nil,
		Mutex:    &sync.Mutex{},
	}
}

func NewWriter(writer ByteWriter) *Writer {
	flushers := make([]Flusher, 0)
	if f, ok := writer.(Flusher); ok {
		flushers = append(flushers, f)
	}
	return &Writer{
		Writer:   writer,
		Flushers: flushers,
		Closer:   nil,
		Mutex:    &sync.Mutex{}}
}

func NewWriterWithClosers(writer ByteWriter, closers ...io.Closer) *Writer {
	flushers := make([]Flusher, 0)
	if f, ok := writer.(Flusher); ok {
		flushers = append(flushers, f)
	}
	for _, c := range closers {
		if f, ok := c.(Flusher); ok {
			flushers = append(flushers, f)
		}
	}
	return &Writer{
		Writer:   writer,
		Flushers: flushers,
		Closer: &Closer{
			Closers: closers,
		},
		Mutex: &sync.Mutex{}}
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

	if w.Flushers != nil {
		for i, f := range w.Flushers {
			err := f.Flush()
			if err != nil {
				return errors.Wrapf(err, "error flushing underlying closer %d", i)
			}
		}
	}

	return nil
}

// FlushSafe flushes any intermediate writer.
// FlushSafe also locks the writer for the duration of flushing using a sync.Mutex.
//  - https://godoc.org/io#Writer
//  - https://godoc.org/sync#Mutex
func (w *Writer) FlushSafe() error {

	/*
		if w.Writer != nil {
			w.Lock()
			var err error
			if flusher, ok := w.Writer.(Flusher); ok {
				err = flusher.Flush()
			}
			w.Unlock()
			if err != nil {
				return errors.Wrap(err, "error flushing underlying writer")
			}
		}*/

	if w.Flushers != nil {
		w.Lock()
		for i, f := range w.Flushers {
			err := f.Flush()
			if err != nil {
				w.Unlock()
				return errors.Wrapf(err, "error flushing underlying closer %d", i)
			}
		}
		w.Unlock()
	}

	return nil
}

// Close flushes the writer and then closes all the underlying io.Closer sequentially.
func (w *Writer) Close() error {
	if w.Closer == nil {
		return nil
	}
	return w.Closer.Close()
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

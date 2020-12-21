// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gzip

import (
	"compress/gzip"
	"fmt"
	"io"
)

type Reader struct {
	Header     gzip.Header
	reader     *gzip.Reader
	underlying io.ReadCloser
}

// Multistream controls whether the reader supports multistream files.
//
// If enabled (the default), the Reader expects the input to be a sequence
// of individually gzipped data streams, each with its own header and
// trailer, ending at EOF. The effect is that the concatenation of a sequence
// of gzipped files is treated as equivalent to the gzip of the concatenation
// of the sequence. This is standard behavior for gzip readers.
//
// Calling Multistream(false) disables this behavior; disabling the behavior
// can be useful when reading file formats that distinguish individual gzip
// data streams or mix gzip data streams with other data streams.
// In this mode, when the Reader reaches the end of the data stream,
// Read returns io.EOF. The underlying reader must implement io.ByteReader
// in order to be left positioned just after the gzip stream.
// To start the next stream, call z.Reset(r) followed by z.Multistream(false).
// If there is no next stream, z.Reset(r) will return io.EOF.
func (z *Reader) Multistream(ok bool) {
	z.reader.Multistream(ok)
}

// Read implements io.Reader, reading uncompressed bytes from its underlying Reader.
func (z *Reader) Read(p []byte) (int, error) {
	return z.reader.Read(p)
}

// Reset discards the Reader z's state and makes it equivalent to the
// result of its original state from NewReader, but reading from r instead.
// This permits reusing a Reader rather than allocating a new one.
func (z *Reader) Reset(reader io.ReadCloser) error {
	z.underlying = nil
	err := z.reader.Reset(reader)
	if err != nil {
		return fmt.Errorf("error resetting reader: %w", err)
	}
	z.underlying = reader
	return nil
}

// Close closes the Reader.
// In order for the GZIP checksum to be verified, the reader must be fully consumed until the io.EOF
// Calls the Close method of the underlying reader, if it implements io.Closer.
func (r *Reader) Close() error {
	err := r.reader.Close()
	if err != nil {
		return fmt.Errorf("error closing reader: %w", err)
	}
	err = r.underlying.Close()
	if err != nil {
		return fmt.Errorf("error closing underlying reader: %w", err)
	}
	return nil
}

// NewReader creates a new Reader reading the given reader.
// If r does not also implement io.ByteReader,
// the decompressor may read more data than necessary from r.
//
// It is the caller's responsibility to call Close on the Reader when done.
//
// The Reader.Header fields will be valid in the Reader returned.
func NewReader(r io.ReadCloser) (*Reader, error) {
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	return &Reader{Header: gr.Header, reader: gr, underlying: r}, nil
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gzip

import (
	"io"

	"compress/gzip"

	"github.com/pkg/errors"
)

type Reader struct {
	*gzip.Reader
	underlying io.Reader
}

// Close closes the Reader.
// In order for the GZIP checksum to be verified, the reader must be fully consumed until the io.EOF
// Calls the Close method of the underlying reader, if it implements io.Closer.
func (r *Reader) Close() error {
	err := r.Reader.Close()
	if err != nil {
		return errors.Wrap(err, "error closing gzip.Reader")
	}
	if c, ok := r.underlying.(io.Closer); ok {
		err = c.Close()
		if err != nil {
			return errors.Wrap(err, "error closing underlying reader")
		}
	}
	return nil
}

// Reset discards the Reader z's state and makes it equivalent to the
// result of its original state from NewReader, but reading from r instead.
// This permits reusing a Reader rather than allocating a new one.
func (r *Reader) Reset(reader io.Reader) error {
	r.underlying = nil
	err := r.Reader.Reset(reader)
	if err != nil {
		return err
	}
	r.underlying = reader
	return nil
}

// NewReader creates a new Reader reading the given reader.
// If r does not also implement io.ByteReader,
// the decompressor may read more data than necessary from r.
//
// It is the caller's responsibility to call Close on the Reader when done.
//
// The Reader.Header fields will be valid in the Reader returned.
func NewReader(r io.Reader) (*Reader, error) {
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	return &Reader{Reader: gr, underlying: r}, nil
}

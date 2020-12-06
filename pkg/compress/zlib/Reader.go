// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package zlib

import (
	"compress/zlib"
	"fmt"
	"io"
)

type Reader struct {
	io.ReadCloser
	underlying io.Reader
}

// Close closes the underlying reader if it implements io.Closer.
func (r *Reader) Close() error {
	if c, ok := r.underlying.(io.Closer); ok {
		err := c.Close()
		if err != nil {
			return fmt.Errorf("error closing underlying reader: %w", err)
		}
	}
	return nil
}

// NewReader creates a new ReadCloser.
// Reads from the returned ReadCloser read and decompress data from r.
// If r does not implement io.ByteReader, the decompressor may read more
// data than necessary from r.
// It is the caller's responsibility to call Close on the ReadCloser when done.
//
// The ReadCloser returned by NewReader also implements Resetter.
func NewReader(r io.Reader) (*Reader, error) {
	zr, err := zlib.NewReader(r)
	if err != nil {
		return nil, err
	}
	return &Reader{ReadCloser: zr, underlying: r}, nil
}

// NewReaderDict is like NewReader but uses a preset dictionary.
// NewReaderDict ignores the dictionary if the compressed data does not refer to it.
// If the compressed data refers to a different dictionary, NewReaderDict returns ErrDictionary.
//
// The ReadCloser returned by NewReaderDict also implements Resetter.
func NewReaderDict(r io.Reader, dict []byte) (*Reader, error) {
	zr, err := zlib.NewReaderDict(r, dict)
	if err != nil {
		return nil, err
	}
	return &Reader{ReadCloser: zr, underlying: r}, nil
}

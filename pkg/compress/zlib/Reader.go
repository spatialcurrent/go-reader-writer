// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package zlib

import (
	"compress/zlib"
	"fmt"
	"io"
)

type ByteReadCloser interface {
	io.ReadCloser
	io.ByteReader
}

type Reader struct {
	reader     io.ReadCloser
	underlying io.ReadCloser
}

// Read implements the io.Reader inferface.
func (r *Reader) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

// Reset resets a ReadCloser returned by NewReader or NewReaderDict
// to switch to a new underlying Reader. This permits reusing a ReadCloser
// instead of allocating a new one.
func (z *Reader) Reset(reader ByteReadCloser, dict []byte) error {
	z.underlying = nil
	if resetter, ok := z.reader.(zlib.Resetter); ok {
		err := resetter.Reset(reader, dict)
		if err != nil {
			return fmt.Errorf("error resetting underlying reader: %w", err)
		}
	} else {
		r, err := zlib.NewReaderDict(reader, dict)
		if err != nil {
			return fmt.Errorf("error recreating reader: %w", err)
		}
		z.reader = r
	}
	z.underlying = reader
	return nil
}

// Close closes the Reader and the underlying io.ReadCloser.
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

// NewReader creates a new ReadCloser.
// Reads from the returned ReadCloser read and decompress data from r.
// If r does not implement io.ByteReader, the decompressor may read more
// data than necessary from r.
// It is the caller's responsibility to call Close on the ReadCloser when done.
//
// The ReadCloser returned by NewReader also implements Resetter.
func NewReader(r ByteReadCloser) (*Reader, error) {
	zr, err := zlib.NewReader(r)
	if err != nil {
		return nil, err
	}
	return &Reader{reader: zr, underlying: r}, nil
}

// NewReaderDict is like NewReader but uses a preset dictionary.
// NewReaderDict ignores the dictionary if the compressed data does not refer to it.
// If the compressed data refers to a different dictionary, NewReaderDict returns ErrDictionary.
//
// The ReadCloser returned by NewReaderDict also implements Resetter.
func NewReaderDict(r ByteReadCloser, dict []byte) (*Reader, error) {
	zr, err := zlib.NewReaderDict(r, dict)
	if err != nil {
		return nil, err
	}
	return &Reader{reader: zr, underlying: r}, nil
}

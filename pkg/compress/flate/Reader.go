// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package flate

import (
	"fmt"
	"io"

	"compress/flate"
)

type ByteReadCloser interface {
	io.ReadCloser
	io.ByteReader
}

type Reader struct {
	reader     io.ReadCloser
	underlying io.ReadCloser
}

func (r *Reader) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

func (r *Reader) Reset(reader ByteReadCloser, dict []byte) error {
	r.underlying = nil
	if resetter, ok := r.reader.(flate.Resetter); ok {
		err := resetter.Reset(reader, dict)
		if err != nil {
			return fmt.Errorf("error resetting underlying reader: %w", err)
		}
	} else {
		r.reader = flate.NewReaderDict(r, dict)
	}
	r.underlying = reader
	return nil
}

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

// NewReader returns a new ReadCloser that can be used
// to read the uncompressed version of r.
// If r does not also implement io.ByteReader,
// the decompressor may read more data than necessary from r.
// It is the caller's responsibility to call Close on the ReadCloser
// when finished reading.
//
// The ReadCloser returned by NewReader also implements Resetter.
func NewReader(r ByteReadCloser) *Reader {
	return &Reader{reader: flate.NewReader(r), underlying: r}
}

// NewReaderDict is like NewReader but initializes the reader
// with a preset dictionary. The returned Reader behaves as if
// the uncompressed data stream started with the given dictionary,
// which has already been read. NewReaderDict is typically used
// to read data compressed by NewWriterDict.
//
// The ReadCloser returned by NewReader also implements Resetter.
func NewReaderDict(r ByteReadCloser, dict []byte) *Reader {
	return &Reader{reader: flate.NewReaderDict(r, dict), underlying: r}
}

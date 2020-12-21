// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package bzip2

import (
	"fmt"
	"io"

	"compress/bzip2"
)

type ByteReadCloser interface {
	io.ReadCloser
	io.ByteReader
}

type Reader struct {
	reader     io.Reader
	underlying io.Closer
}

func (r *Reader) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

func (r *Reader) Close() error {
	err := r.underlying.Close()
	if err != nil {
		return fmt.Errorf("error closing underlying reader: %w", err)
	}
	return nil
}

// NewReader returns an io.Reader which decompresses bzip2 data from r.
// If r does not also implement io.ByteReader, the decompressor may read more data than necessary from r.
func NewReader(r ByteReadCloser) *Reader {
	return &Reader{reader: bzip2.NewReader(r), underlying: r}
}

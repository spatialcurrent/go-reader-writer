// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package snappy

import (
	"io"

	"github.com/golang/snappy"
)

type Reader struct {
	reader     *snappy.Reader
	underlying io.ReadCloser
}

func (r *Reader) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

func (r *Reader) Reset(reader io.ReadCloser) {
	r.reader.Reset(reader)
	r.underlying = reader
}

func (r *Reader) Close() error {
	return r.underlying.Close()
}

// NewReader returns a new Reader that decompresses from r, using the framing
// format described at
// https://github.com/google/snappy/blob/master/framing_format.txt
func NewReader(r io.ReadCloser) *Reader {
	return &Reader{reader: snappy.NewReader(r), underlying: r}
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package flate

import (
	"io"

	"compress/flate"

	"github.com/pkg/errors"
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
			return errors.Wrap(err, "error closing underlying reader")
		}
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
func NewReader(r io.Reader) *Reader {
	return &Reader{ReadCloser: flate.NewReader(r), underlying: r}
}

// NewReaderDict is like NewReader but initializes the reader
// with a preset dictionary. The returned Reader behaves as if
// the uncompressed data stream started with the given dictionary,
// which has already been read. NewReaderDict is typically used
// to read data compressed by NewWriterDict.
//
// The ReadCloser returned by NewReader also implements Resetter.
func NewReaderDict(r io.Reader, dict []byte) *Reader {
	return &Reader{ReadCloser: flate.NewReaderDict(r, dict), underlying: r}
}

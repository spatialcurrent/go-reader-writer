// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package zip

import (
	"archive/zip"
	"fmt"
	"io"
)

type Reader struct {
	*zip.Reader
	underlying io.ReaderAt
}

// Close closes the Zip file, rendering it unusable for I/O.
func (r *Reader) Close() error {
	if r.underlying != nil {
		if c, ok := r.underlying.(io.Closer); ok {
			err := c.Close()
			if err != nil {
				return fmt.Errorf("error closing underlying reader: %w", err)
			}
		}
	}
	return nil
}

// NewReader returns a new Reader reading from r, which is assumed to have the given size in bytes.
func NewReader(r io.ReaderAt, size int64) (*Reader, error) {
	zr, err := zip.NewReader(r, size)
	if err != nil {
		return nil, err
	}
	return &Reader{Reader: zr, underlying: r}, nil
}

// OpenReader will open the Zip file specified by name and return a ReadCloser.
func OpenReader(name string) (*zip.ReadCloser, error) {
	zr, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}
	return zr, nil
}

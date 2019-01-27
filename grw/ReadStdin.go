// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"compress/bzip2"
	"compress/gzip"
	"os"
)

import (
	"github.com/golang/snappy"
	"github.com/pkg/errors"
)

// ReadStdin returns a ByteReadCloser for a file with a given compression.
// alg may be "bzip2", "gzip", "snappy", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func ReadStdin(alg string, cache bool) (ByteReadCloser, error) {
	switch alg {
	case "bzip2":
		br := bzip2.NewReader(os.Stdin)
		if cache {
			return &Cache{
				Reader:  &Reader{Reader: bufio.NewReader(br)},
				Content: &[]byte{},
			}, nil
		}
		return &Reader{Reader: bufio.NewReader(br)}, nil
	case "gzip":
		gr, err := gzip.NewReader(bufio.NewReader(os.Stdin))
		if err != nil {
			return nil, errors.Wrap(err, "Error creating gzip reader for stdin")
		}
		if cache {
			return &Cache{
				Reader:  &Reader{Reader: bufio.NewReader(gr), Closer: gr},
				Content: &[]byte{},
			}, nil
		}
		return &Reader{Reader: bufio.NewReader(gr), Closer: gr}, nil
	case "snappy":
		sr := snappy.NewReader(bufio.NewReader(os.Stdin))
		if cache {
			return NewCache(&Reader{Reader: bufio.NewReader(sr)}), nil
		}
		return &Reader{Reader: bufio.NewReader(sr)}, nil
	case "none", "":
		br := bufio.NewReader(os.Stdin)
		if cache {
			return NewCache(&Reader{Reader: br}), nil
		}
		return &Reader{Reader: br}, nil
	}

	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

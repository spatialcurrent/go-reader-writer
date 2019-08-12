// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"io"
)

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

// ReadFromFilePath returns a ByteReader for a file with a given compression.
// alg may be "snappy", "gzip", or "none."
//
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func ReadFromFilePath(path string, alg string, bufferSize int) (ByteReadCloser, error) {

	if len(path) == 0 {
		return nil, ErrPathMissing
	}

	pathExpanded, err := homedir.Expand(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error expanding file path %q", path)
	}

	switch alg {
	case AlgorithmBzip2, AlgorithmGzip, AlgorithmSnappy, AlgorithmNone, "":
		f, err := OpenFile(pathExpanded)
		if err != nil {
			return nil, errors.Wrapf(err, "error opening file at path %q", pathExpanded)
		}
		r, closers, err := WrapReader(f, []io.Closer{f}, alg, bufferSize)
		if err != nil {
			return nil, errors.Wrapf(err, "error wrapping reader for file at path %q", pathExpanded)
		}
		return &Reader{Reader: r, Closers: closers}, nil
	case AlgorithmZip:
		return ReadZipFile(pathExpanded)
	}

	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

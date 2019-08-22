// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"io"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

// ReadFromFilePath opens a ByteReadCloser for the given path, compression algorithm, and buffer size.
// ReadFromFilePath returns the ByteReadCloser and error, if any.
//
// ReadFromFilePath returns an error if the path cannot be expanded,
// the file references by the path cannot be opened, or
// the compression algorithm is invalid.
//
func ReadFromFilePath(path string, alg string, bufferSize int) (ByteReadCloser, error) {

	if len(path) == 0 {
		return nil, ErrPathMissing
	}

	pathExpanded, err := homedir.Expand(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error expanding file path %q", path)
	}

	pathAbsolute, err := filepath.Abs(pathExpanded)
	if err != nil {
		return nil, errors.Wrapf(err, "error resolving file path %q", pathAbsolute)
	}

	switch alg {
	case AlgorithmBzip2, AlgorithmGzip, AlgorithmSnappy, AlgorithmNone, "":
		f, err := OpenFile(pathAbsolute)
		if err != nil {
			return nil, errors.Wrapf(err, "error opening file at path %q", pathAbsolute)
		}
		r, closers, err := WrapReader(f, []io.Closer{f}, alg, bufferSize)
		if err != nil {
			return nil, errors.Wrapf(err, "error wrapping reader for file at path %q", pathAbsolute)
		}
		return &Reader{Reader: r, Closer: &Closer{Closers: closers}}, nil
	case AlgorithmZip:
		return ReadZipFile(pathAbsolute)
	}

	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

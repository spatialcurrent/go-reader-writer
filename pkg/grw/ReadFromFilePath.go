// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

/*

import (
	"fmt"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/spatialcurrent/go-reader-writer/pkg/os"
)

// ReadFromFilePathInput contains the input parameters ReadFromFilePath.
type ReadFromFilePathInput struct {
	Path       string // the path to the file
	Alg        string // the compression algorithm used
	Dict       []byte // the dictionary for the compression algorithm, if applicable
	BufferSize int    // the buffer size for the underlying reader
}

// ReadFromFilePath opens a ByteReadCloser for the given path, compression algorithm, and buffer size.
// ReadFromFilePath returns the ByteReadCloser and error, if any.
//
// ReadFromFilePath returns an error if the path cannot be expanded,
// the file references by the path cannot be opened, or
// the compression algorithm is invalid.
//
func ReadFromFilePath(input *ReadFromFilePathInput) (*Reader, error) {

	if len(input.Path) == 0 {
		return nil, ErrPathMissing
	}

	pathExpanded, err := homedir.Expand(input.Path)
	if err != nil {
		return nil, fmt.Errorf("error expanding file path %q: %w", input.Path, err)
	}

	pathCleaned := filepath.Clean(pathExpanded)

	switch input.Alg {
	case AlgorithmBzip2, AlgorithmFlate, AlgorithmGzip, AlgorithmNone, AlgorithmSnappy, AlgorithmZlib, "":
		f, err := os.OpenFile(pathCleaned)
		if err != nil {
			return nil, fmt.Errorf("error opening file at path %q: %w", pathCleaned, err)
		}
		r, err := WrapReader(f, input.Alg, input.Dict, input.BufferSize)
		if err != nil {
			return nil, fmt.Errorf("error wrapping reader for file at path %q: %w", pathCleaned, err)
		}
		return &Reader{Reader: r}, nil
	case AlgorithmZip:
		return ReadZipFile(pathCleaned)
	}

	return nil, &ErrUnknownAlgorithm{Algorithm: input.Alg}
}
*/

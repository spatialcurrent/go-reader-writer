// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
	"io"
	"path/filepath"

	pkgalg "github.com/spatialcurrent/go-reader-writer/pkg/alg"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/flate"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/gzip"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/snappy"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/zlib"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
)

// WriteToFileSystemInput contains the input parameters for WriteToFileSystem.
type WriteToFileSystemInput struct {
	Alg        string // compression algorithm
	BufferSize int    // buffer size
	Dict       []byte // compression dictionary
	Flag       int    // flag for file descriptor
	Path       string // path to write to
	Parents    bool   // automatically create parent directories as necessary
}

// WriteToFileSystem returns a ByteWriteCloser for a file with a given compression.
// alg may be "snappy", "gzip", or "none."
//
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func WriteToFileSystem(input *WriteToFileSystemInput) (io.WriteCloser, error) {

	if input.Parents {
		err := os.MkdirAll(filepath.Dir(input.Path), 0770)
		if err != nil {
			return nil, fmt.Errorf("error creating parent directories: %w", err)
		}
	}

	/*
		fileInfo, err := os.Stat(input.Path)
		if err != nil {
			return nil, fmt.Errorf("error stating file at path %q: %w", input.Path, err)
		}

		if fileInfo.Mode()&os.ModeDevice != 0 {
			close = false
		}
	*/

	switch input.Alg {
	case pkgalg.AlgorithmBzip2:
		return nil, &ErrWriterNotImplemented{Algorithm: input.Alg}
	case pkgalg.AlgorithmFlate:
		return flate.WriteFile(input.Path, input.BufferSize)
	case pkgalg.AlgorithmGzip:
		return gzip.WriteFile(input.Path, input.BufferSize)
	case pkgalg.AlgorithmSnappy:
		return snappy.WriteFile(input.Path, input.BufferSize)
	case pkgalg.AlgorithmZip:
		return nil, &ErrWriterNotImplemented{Algorithm: input.Alg}
	case pkgalg.AlgorithmZlib:
		return zlib.WriteFile(input.Path, input.Dict, input.BufferSize)
	case pkgalg.AlgorithmNone:
		return WriteFile(input.Path, input.Dict, input.BufferSize)
	}

	return nil, &pkgalg.ErrUnknownAlgorithm{Algorithm: input.Alg}
}

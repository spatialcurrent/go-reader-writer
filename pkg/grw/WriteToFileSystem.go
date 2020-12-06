// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
	"path/filepath"

	"github.com/spatialcurrent/go-reader-writer/pkg/os"
)

// WriteToFileSystemInput contains the input parameters for WriteToFileSystem.
type WriteToFileSystemInput struct {
	Path    string // path to write to
	Alg     string // compression algorithm
	Dict    []byte // compression dictionary
	Flag    int    // flag for file descriptor
	Parents bool   // automatically create parent directories as necessary
}

// WriteToFileSystem returns a ByteWriteCloser for a file with a given compression.
// alg may be "snappy", "gzip", or "none."
//
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func WriteToFileSystem(input *WriteToFileSystemInput) (*Writer, error) {

	if input.Parents {
		err := os.MkdirAll(filepath.Dir(input.Path), 0770)
		if err != nil {
			return nil, fmt.Errorf("error creating parent directories: %w", err)
		}
	}

	switch input.Alg {
	case AlgorithmBzip2:
		return nil, &ErrWriterNotImplemented{Algorithm: "bzip2"}
	case AlgorithmFlate:
		return WriteFlateFile(input.Path, input.Dict, input.Flag)
	case AlgorithmGzip:
		return WriteGzipFile(input.Path, input.Flag)
	case AlgorithmSnappy:
		return WriteSnappyFile(input.Path, input.Flag)
	case AlgorithmZip:
		return nil, &ErrWriterNotImplemented{Algorithm: "zip"}
	case AlgorithmZlib:
		return WriteZlibFile(input.Path, input.Dict, input.Flag)
	case AlgorithmNone, "":
		return WriteLocalFile(input.Path, input.Flag)
	}
	return nil, &ErrUnknownAlgorithm{Algorithm: input.Alg}
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mitchellh/go-homedir"

	pkgalg "github.com/spatialcurrent/go-reader-writer/pkg/alg"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

type WriteToResourceInput struct {
	Alg        string // compression algorithm
	Append     bool   // append to output resource
	BufferSize int    // buffer size
	Dict       []byte // compression dictionary
	Parents    bool   // automatically create parent directories as necessary
	S3Client   *s3.S3 // AWS S3 Client
	URI        string // uri to write to
}

type WriteToResourceOutput struct {
	Writer io.WriteCloser
}

// WriteToResource returns a ByteWriteCloser and error, if any.
func WriteToResource(input *WriteToResourceInput) (*WriteToResourceOutput, error) {

	if input.URI == "-" {
		w, err := WrapWriter(os.Stdout, input.Alg, input.Dict, 0)
		if err != nil {
			return nil, fmt.Errorf("error wrapping device %q: %w", input.URI, err)
		}
		return &WriteToResourceOutput{Writer: w}, nil
	}

	scheme, path := splitter.SplitUri(input.URI)
	switch scheme {
	case "none", "":

		pathExpanded, err := homedir.Expand(path)
		if err != nil {
			return nil, fmt.Errorf("error expanding resource file path %q: %w", path, err)
		}

		flag := 0
		if input.Append {
			flag = os.O_APPEND | os.O_CREATE | os.O_WRONLY
		} else {
			flag = os.O_CREATE | os.O_WRONLY
		}

		if input.Parents {
			err = os.MkdirAll(filepath.Dir(pathExpanded), 0770)
			if err != nil {
				return nil, fmt.Errorf("error creating parent directories: %w", err)
			}
		}

		w, err := WriteToFileSystem(&WriteToFileSystemInput{
			Alg:        input.Alg,
			BufferSize: input.BufferSize,
			Dict:       input.Dict,
			Flag:       flag,
			Parents:    false,
			Path:       pathExpanded,
		})
		return &WriteToResourceOutput{Writer: w}, nil
	}

	return nil, &pkgalg.ErrUnknownAlgorithm{Algorithm: input.Alg}
}

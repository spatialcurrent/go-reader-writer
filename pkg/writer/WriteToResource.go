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

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mitchellh/go-homedir"

	"github.com/spatialcurrent/go-reader-writer/pkg/os"
	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

type WriteToResourceInput struct {
	Uri      string // uri to write to
	Alg      string // compression algorithm
	Dict     []byte // compression dictionary
	Append   bool   // append to output resource
	Parents  bool   // automatically create parent directories as necessary
	S3Client *s3.S3 // AWS S3 Client
}

// WriteToResource returns a ByteWriteCloser and error, if any.
func WriteToResource(input *WriteToResourceInput) (*Writer, error) {

	if outputDevice := os.OpenDevice(input.Uri); outputDevice != nil {
		w, err := WrapWriter(outputDevice, input.Alg, input.Dict)
		if err != nil {
			return nil, fmt.Errorf("error wrapping device %q: %w", input.Uri, err)
		}
		return w, nil
	}

	scheme, path := splitter.SplitUri(input.Uri)
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

		return WriteToFileSystem(&WriteToFileSystemInput{
			Path:    pathExpanded,
			Alg:     input.Alg,
			Dict:    input.Dict,
			Flag:    flag,
			Parents: false,
		})
	}

	return nil, &ErrUnknownAlgorithm{Algorithm: input.Alg}
}

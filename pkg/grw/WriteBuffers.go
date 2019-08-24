// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bytes"
	"io"
	"path/filepath"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

func WriteBuffers(buffers map[string]Buffer, alg string, append bool, mkdirs bool, s3Client *s3.S3) error {

	for uri, buffer := range buffers {

		data := buffer.Bytes()

		scheme, path := splitter.SplitUri(uri)

		// If output is a file, then create parent directories.
		if scheme == "" || scheme == "file" {
			if mkdirs {
				err := Mkdirs(filepath.Dir(path))
				if err != nil {
					return errors.Wrapf(err, "error creating parent directories for path %q", uri)
				}
			}
		}

		writer, err := WriteToResource(uri, alg, append, s3Client)
		if err != nil {
			return errors.Wrapf(err, "error opening output file for path %q", uri)
		}

		_, err = io.Copy(writer, bytes.NewReader(data))
		if err != nil {
			return errors.Wrapf(err, "error writing to output file for uri %q", uri)
		}

		err = writer.Flush()
		if err != nil {
			return errors.Wrapf(err, "error flushing to output file for uri %q", uri)
		}

		err = writer.Close()
		if err != nil {
			return errors.Wrapf(err, "error closing output file for uri %q", uri)
		}
	}

	return nil
}

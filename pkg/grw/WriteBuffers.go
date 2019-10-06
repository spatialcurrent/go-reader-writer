// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-reader-writer/pkg/io"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

type WriteBuffersInput struct {
	Buffers    map[string]io.Buffer
	Algorithm  string
	Dictionary []byte
	Overwrite  bool
	Append     bool
	Mkdirs     bool
	S3Client   *s3.S3
}

// WriteBuffers writes a map of buffers to the resources defined by the keys.
// alg is the compression algorithm.
// If the buffer already includes compressed data, then use "" or "none" as alg.
// If append is true, then append to existing files.
// If mkdirs is true, then parent directories are created on-demand.
func WriteBuffers(input *WriteBuffersInput) error {

	for uri, buffer := range input.Buffers {

		scheme, path := splitter.SplitUri(uri)

		if scheme == "" || scheme == "file" {
			if (!input.Overwrite) && (!input.Append) {
				exists, _, err := os.Stat(path)
				if err != nil {
					return errors.Wrapf(err, "error statting uri %q", uri)
				}
				if exists {
					return fmt.Errorf("file already exists at uri %q and neither append or overwrite is set", uri)
				}
			}
		}

		writer, err := WriteToResource(&WriteToResourceInput{
			Uri:      uri,
			Alg:      input.Algorithm,
			Dict:     input.Dictionary,
			Append:   input.Append,
			Parents:  input.Mkdirs,
			S3Client: input.S3Client,
		})
		if err != nil {
			return errors.Wrapf(err, "error opening output file for path %q", uri)
		}

		_, err = io.Copy(writer, buffer)
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

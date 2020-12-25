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
					return fmt.Errorf("error statting uri %q: %w", uri, err)
				}
				if exists {
					return fmt.Errorf("file already exists at uri %q and neither append or overwrite is set", uri)
				}
			}
		}

		writeToResourceOutput, err := WriteToResource(&WriteToResourceInput{
			URI:      uri,
			Alg:      input.Algorithm,
			Dict:     input.Dictionary,
			Append:   input.Append,
			Parents:  input.Mkdirs,
			S3Client: input.S3Client,
		})
		if err != nil {
			return fmt.Errorf("error opening output file for path %q: %w", uri, err)
		}
		writer := writeToResourceOutput.Writer

		_, err = io.Copy(writer, buffer)
		if err != nil {
			return fmt.Errorf("error writing to output file for uri %q: %w", uri, err)
		}

		if flusher, ok := writer.(interface{ Flush() error }); ok {
			err = flusher.Flush()
			if err != nil {
				return fmt.Errorf("error flushing to output file for uri %q: %w", uri, err)
			}
		}

		err = writer.Close()
		if err != nil {
			return fmt.Errorf("error closing output file for uri %q: %w", uri, err)
		}
	}

	return nil
}

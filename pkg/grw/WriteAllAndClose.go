// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

// WriteAllAndCloseInput contains the input parameters for WriteAllAndClose.
type WriteAllAndCloseInput struct {
	ACL        string // ACL for objects written to AWS S3
	BufferSize int    // buffer size
	Bytes      []byte // the content to write
	URI        string // uri to write to
	Alg        string // compression algorithm
	Dict       []byte // compression dictionary
	Append     bool   // append to output resource
	Parents    bool   // automatically create parent directories as necessary
	S3Client   *s3.S3 // AWS S3 Client
}

// WriteAllAndClose writes the bytes to the resource indicated by the uri given, flushes, and closes the resource.
func WriteAllAndClose(input *WriteAllAndCloseInput) error {
	scheme, path := splitter.SplitURI(input.URI)
	switch scheme {
	case "s3":
		i := strings.Index(path, "/")
		if i == -1 {
			return errors.New("s3 path missing bucket")
		}
		err := UploadS3Object(&UploadS3ObjectInput{
			ACL:    input.ACL,
			Bucket: path[0:i],
			Client: input.S3Client,
			Key:    path[i+1:],
			Object: bytes.NewBuffer(input.Bytes),
		})
		if err != nil {
			return fmt.Errorf("error uploading new version of catalog to S3: %w", err)
		}
		return nil
	case "file", "":
		writeToResourceOutput, err := WriteToResource(&WriteToResourceInput{
			ACL:        input.ACL,
			Alg:        input.Alg,
			Append:     input.Append,
			BufferSize: input.BufferSize,
			Dict:       input.Dict,
			Parents:    input.Parents,
			S3Client:   input.S3Client,
			URI:        input.URI,
		})
		if err != nil {
			return fmt.Errorf("error opening resource at uri %q: %w", input.URI, err)
		}
		w := writeToResourceOutput.Writer
		_, err = w.Write(input.Bytes)
		if err != nil {
			return fmt.Errorf("error writing to resource at uri %q: %w", input.URI, err)
		}
		if flusher, ok := w.(interface{ Flush() error }); ok {
			err = flusher.Flush()
			if err != nil {
				return fmt.Errorf("error flushing to resource at uri %q: %w", input.URI, err)
			}
		}
		err = w.Close()
		if err != nil {
			return fmt.Errorf("error closing resource at uri %q: %w", input.URI, err)
		}
	}
	return nil
}

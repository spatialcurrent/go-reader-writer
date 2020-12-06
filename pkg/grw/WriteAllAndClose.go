// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
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
	Bytes    []byte // the content to write
	Uri      string // uri to write to
	Alg      string // compression algorithm
	Dict     []byte // compression dictionary
	Append   bool   // append to output resource
	Parents  bool   // automatically create parent directories as necessary
	S3Client *s3.S3 // AWS S3 Client
}

// WriteAllAndClose writes the bytes to the resource indicated by the uri given, flushes, and closes the resource.
func WriteAllAndClose(input *WriteAllAndCloseInput) error {
	scheme, path := splitter.SplitUri(input.Uri)
	switch scheme {
	case "s3":
		i := strings.Index(path, "/")
		if i == -1 {
			return errors.New("s3 path missing bucket")
		}
		err := UploadS3Object(path[0:i], path[i+1:], bytes.NewBuffer(input.Bytes), input.S3Client)
		if err != nil {
			return fmt.Errorf("error uploading new version of catalog to S3: %w", err)
		}
		return nil
	case "file", "":
		w, err := WriteToResource(&WriteToResourceInput{
			Uri:      input.Uri,
			Alg:      input.Alg,
			Dict:     input.Dict,
			Append:   input.Append,
			Parents:  input.Parents,
			S3Client: input.S3Client,
		})
		if err != nil {
			return fmt.Errorf("error opening resource at uri %q: %w", input.Uri, err)
		}
		_, err = w.Write(input.Bytes)
		if err != nil {
			return fmt.Errorf("error writing to resource at uri %q: %w", input.Uri, err)
		}
		err = w.Flush()
		if err != nil {
			return fmt.Errorf("error flushing to resource at uri %q: %w", input.Uri, err)
		}
		err = w.Close()
		if err != nil {
			return fmt.Errorf("error closing resource at uri %q: %w", input.Uri, err)
		}
	}
	return nil
}

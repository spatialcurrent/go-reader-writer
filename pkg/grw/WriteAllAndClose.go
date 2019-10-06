// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
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
	w, err := WriteToResource(&WriteToResourceInput{
		Uri:      input.Uri,
		Alg:      input.Alg,
		Dict:     input.Dict,
		Append:   input.Append,
		Parents:  input.Parents,
		S3Client: input.S3Client,
	})
	if err != nil {
		return errors.Wrapf(err, "error opening resource at uri %q", input.Uri)
	}

	_, err = w.Write(input.Bytes)
	if err != nil {
		return errors.Wrapf(err, "error writing to resource at uri %q", input.Uri)
	}

	err = w.Flush()
	if err != nil {
		return errors.Wrapf(err, "error flushing to resource at uri %q", input.Uri)
	}

	err = w.Close()
	if err != nil {
		return errors.Wrapf(err, "error closing resource at uri %q", input.Uri)
	}

	return nil
}

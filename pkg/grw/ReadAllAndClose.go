// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

/*
import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
)

type ReadAllAndCloseInput struct {
	Uri        string
	Alg        string
	Dict       []byte
	BufferSize int
	S3Client   *s3.S3
}

func ReadAllAndClose(input *ReadAllAndCloseInput) ([]byte, error) {
	r, _, err := ReadFromResource(&ReadFromResourceInput{
		Uri:        input.URI,
		Alg:        input.Alg,
		Dict:       input.Dict,
		BufferSize: input.BufferSize,
		S3Client:   input.S3Client,
	})
	if err != nil {
		return make([]byte, 0), fmt.Errorf("error opening resource at uri %q: %w", input.URI, err)
	}
	b, err := r.ReadAllAndClose()
	if err != nil {
		return make([]byte, 0), fmt.Errorf("error reading from resource at uri %q: %w", input.URI, err)
	}
	return b, nil
}
*/

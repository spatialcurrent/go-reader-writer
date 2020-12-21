// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

type ReadFromResourceInput struct {
	Uri        string // uri to read from
	Alg        string // compression algorithm
	Dict       []byte // compression dictionary
	BufferSize int    // input reader buffer size
	S3Client   *s3.S3 // AWS S3 Client
}

type ReadFromResourceOutput struct {
	Reader   *Reader
	Metadata *Metadata
}

func ReadFromResource(input *ReadFromResourceInput) (*ReadFromResourceOutput, error) {
	b := NewBuilder().Uri(input.Uri)
	if len(input.Alg) > 0 {
		b = b.Algorithm(input.Alg)
	}
	if len(input.Dict) > 0 {
		b = b.Dictionary(input.Dict)
	}
	if input.BufferSize > 0 {
		b = b.BufferSize(input.BufferSize)
	}
	if input.S3Client != nil {
		b = b.S3Client(input.S3Client)
	}
	return b.Open()
}

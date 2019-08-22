// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"github.com/aws/aws-sdk-go/service/s3"
	//"strings"
)

//"github.com/pkg/errors"

//"github.com/spatialcurrent/go-reader-writer/pkg/splitter"

type Builder struct {
	uri        string
	alg        string
	bufferSize int
	s3Client   *s3.S3
}

func NewBuilder() *Builder {
	return &Builder{
		uri:        "",
		alg:        "",
		bufferSize: 4096,
		s3Client:   nil,
	}
}

func (b *Builder) Uri(uri string) *Builder {
	return &Builder{
		uri:        uri,
		alg:        b.alg,
		bufferSize: b.bufferSize,
		s3Client:   b.s3Client,
	}
}

func (b *Builder) Algorithm(alg string) *Builder {
	return &Builder{
		uri:        b.uri,
		alg:        alg,
		bufferSize: b.bufferSize,
		s3Client:   b.s3Client,
	}
}

func (b *Builder) BufferSize(bufferSize int) *Builder {
	return &Builder{
		uri:        b.uri,
		alg:        b.alg,
		bufferSize: bufferSize,
		s3Client:   b.s3Client,
	}
}

func (b *Builder) S3Client(s3Client *s3.S3) *Builder {
	return &Builder{
		uri:        b.uri,
		alg:        b.alg,
		bufferSize: b.bufferSize,
		s3Client:   s3Client,
	}
}

func (b *Builder) wrapNoMetadata(brc ByteReadCloser, err error) (ByteReadCloser, *Metadata, error) {
	return brc, nil, err
}

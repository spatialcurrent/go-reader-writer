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

type Builder struct {
	uri        string
	alg        string
	dict       []byte
	bufferSize int
	s3Client   *s3.S3
}

func NewBuilder() *Builder {
	return &Builder{
		uri:        "",
		alg:        "",
		dict:       make([]byte, 0),
		bufferSize: 4096,
		s3Client:   nil,
	}
}

func (b *Builder) Uri(uri string) *Builder {
	return &Builder{
		uri:        uri,
		alg:        b.alg,
		dict:       b.dict,
		bufferSize: b.bufferSize,
		s3Client:   b.s3Client,
	}
}

func (b *Builder) Algorithm(alg string) *Builder {
	return &Builder{
		uri:        b.uri,
		alg:        alg,
		dict:       b.dict,
		bufferSize: b.bufferSize,
		s3Client:   b.s3Client,
	}
}

func (b *Builder) Dictionary(dict []byte) *Builder {
	return &Builder{
		uri:        b.uri,
		alg:        b.alg,
		dict:       dict,
		bufferSize: b.bufferSize,
		s3Client:   b.s3Client,
	}
}

func (b *Builder) BufferSize(bufferSize int) *Builder {
	return &Builder{
		uri:        b.uri,
		alg:        b.alg,
		dict:       b.dict,
		bufferSize: bufferSize,
		s3Client:   b.s3Client,
	}
}

func (b *Builder) S3Client(s3Client *s3.S3) *Builder {
	return &Builder{
		uri:        b.uri,
		alg:        b.alg,
		dict:       b.dict,
		bufferSize: b.bufferSize,
		s3Client:   s3Client,
	}
}

func (b *Builder) wrapNoMetadata(r *Reader, err error) (*Reader, *Metadata, error) {
	return r, nil, err
}

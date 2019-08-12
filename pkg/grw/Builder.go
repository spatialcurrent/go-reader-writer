// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"strings"
)

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

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

func (b *Builder) Open() (ByteReadCloser, *Metadata, error) {

	if b.uri == "stdin" {
		brc, err := ReadStdin(b.uri)
		return brc, nil, err
	}

	scheme, path := SplitUri(b.uri)

	switch scheme {
	case SchemeFtp, "ftps", "sftp":
		return b.wrapNoMetadata(ReadFTPFile(b.uri, b.alg, b.bufferSize))
	case SchemeHttp, SchemeHttps:
		return ReadHTTPFile(b.uri, b.alg, b.bufferSize)
	case SchemeS3:
		i := strings.Index(path, "/")
		if i == -1 {
			return &Reader{}, nil, errors.New("path missing bucket")
		}
		return ReadS3Object(path[0:i], path[i+1:], b.alg, b.bufferSize, b.s3Client)
	case SchemeFile, "none", "":
		return b.wrapNoMetadata(ReadFromFilePath(path, b.alg, b.bufferSize))
	}
	return nil, nil, &ErrUnknownAlgorithm{Algorithm: b.alg}
}

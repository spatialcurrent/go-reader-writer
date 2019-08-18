// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// +build !js

package grw

import (
	"strings"
)

import (
	//"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

import (
	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

func (b *Builder) Open() (ByteReadCloser, *Metadata, error) {

	if b.uri == "stdin" {
		brc, err := ReadStdin(b.uri)
		return brc, nil, err
	}

	scheme, path := splitter.SplitUri(b.uri)

	switch scheme {
	case SchemeFtp, "ftps", "sftp":
		return b.wrapNoMetadata(ReadFTPFile(b.uri, b.alg, b.bufferSize))
	case SchemeHTTP, SchemeHTTPS:
		return ReadHTTPFile(b.uri, b.alg, b.bufferSize)
	case SchemeS3:
		i := strings.Index(path, "/")
		if i == -1 {
			return nil, nil, errors.New("path missing bucket")
		}
		return ReadS3Object(path[0:i], path[i+1:], b.alg, b.bufferSize, b.s3Client)
	case SchemeFile, "none", "":
		return b.wrapNoMetadata(ReadFromFilePath(path, b.alg, b.bufferSize))
	}
	return nil, nil, &ErrUnknownScheme{Scheme: scheme}
}

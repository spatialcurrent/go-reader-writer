// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// +build js

package grw

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

func (b *Builder) Open() (*Reader, *Metadata, error) {

	if b.uri == "stdin" {
		brc, err := ReadStdin(b.alg, b.dict, b.bufferSize)
		return brc, nil, err
	}

	scheme, path := splitter.SplitUri(b.uri)

	switch scheme {
	case SchemeHTTP, SchemeHTTPS:
		return ReadHTTPFile(b.uri, b.alg, b.dict, b.bufferSize)
	case SchemeS3:
		i := strings.Index(path, "/")
		if i == -1 {
			return nil, nil, errors.New("path missing bucket")
		}
		return ReadS3Object(path[0:i], path[i+1:], b.alg, b.dict, b.bufferSize, b.s3Client)
	case SchemeFile, "none", "":
		return b.wrapNoMetadata(ReadFromFilePath(path, b.alg, b.dict, b.bufferSize))
	}
	return nil, nil, &ErrUnknownScheme{Scheme: scheme}
}

// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"strings"
)

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

func ReadFromResource(uri string, alg string, buffer_size int, cache bool, s3_client *s3.S3) (ByteReadCloser, *Metadata, error) {

	if uri == "stdin" {
		brc, err := ReadStdin(alg, cache)
		return brc, nil, err
	}

	scheme, path := SplitUri(uri)
	switch scheme {
	case "http", "https":
		return ReadHTTPFile(uri, alg, cache)
	case "s3":
		i := strings.Index(path, "/")
		if i == -1 {
			return &Reader{}, nil, errors.New("path missing bucket")
		}
		return ReadS3Object(path[0:i], path[i+1:], alg, cache, s3_client)
	case "none", "":
		pathExpanded, err := homedir.Expand(path)
		if err != nil {
			return nil, nil, errors.Wrap(err, "Error expanding resource file path "+path)
		}
		brc, err := ReadFromFilePath(pathExpanded, alg, cache, buffer_size)
		return brc, nil, err
	}

	return nil, nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

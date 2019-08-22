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

func ReadAllAndClose(uri string, alg string, s3Client *s3.S3) ([]byte, error) {
	r, _, err := ReadFromResource(uri, alg, 4096, s3Client)
	if err != nil {
		return make([]byte, 0), errors.Wrapf(err, "error opening resource at uri %q", uri)
	}
	b, err := r.ReadAllAndClose()
	if err != nil {
		return make([]byte, 0), errors.Wrapf(err, "error reading from resource at uri %q", uri)
	}
	return b, nil
}

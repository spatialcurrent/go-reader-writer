// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
)

type Metadata struct {
	ContentType   string
	LastModified  *time.Time
	ContentLength int64
	Header        map[string][]string
}

func NewMetadataFromHeader(header map[string][]string) *Metadata {

	contentType := ""
	if contentTypes, ok := header["Content-Type"]; ok && len(contentTypes) > 0 {
		contentType = contentTypes[0]
	}

	return &Metadata{ContentType: contentType, Header: header}
}

func NewMetadataFromS3(output *s3.GetObjectOutput) *Metadata {
	return &Metadata{ContentType: *output.ContentType, LastModified: output.LastModified, ContentLength: *output.ContentLength}
}

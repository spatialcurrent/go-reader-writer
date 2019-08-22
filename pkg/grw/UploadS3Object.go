// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

// UploadS3Object uploads an object to S3.
// alg may be "bzip2", "gzip", "snappy", "zip", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func UploadS3Object(bucket string, key string, object io.Reader, s3_client *s3.S3) error {

	uploader := s3manager.NewUploaderWithClient(s3_client)

	uploadInput := &s3manager.UploadInput{
		Body:   object,
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	_, err := uploader.Upload(uploadInput)
	if err != nil {
		return errors.Wrap(err, "error uploading data to AWS S3")
	}

	return nil

}

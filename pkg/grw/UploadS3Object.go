// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"errors"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UploadS3ObjectInput struct {
	ACL    string
	Bucket string
	Key    string
	Object io.Reader
	Client *s3.S3
}

// UploadS3Object uploads an object to S3.
// alg may be "bzip2", "gzip", "snappy", "zip", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func UploadS3Object(input *UploadS3ObjectInput) error {

	if input == nil {
		return errors.New("input is nil")
	}

	if len(input.Bucket) == 0 {
		return errors.New("invalid input: bucket is missing")
	}

	if len(input.Key) == 0 {
		return errors.New("invalid input: key is missing")
	}

	if input.Client == nil {
		return errors.New("invalid input: client is nil")
	}

	if input.Object == nil {
		return errors.New("invalid input: object is nil")
	}

	uploader := s3manager.NewUploaderWithClient(input.Client)

	uploadInput := &s3manager.UploadInput{
		Body:   input.Object,
		Bucket: aws.String(input.Bucket),
		Key:    aws.String(input.Key),
	}

	if len(input.ACL) > 0 {
		uploadInput.ACL = aws.String(input.ACL)
	}

	_, err := uploader.Upload(uploadInput)
	if err != nil {
		return fmt.Errorf("error uploading data to AWS S3: %w", err)
	}

	return nil

}

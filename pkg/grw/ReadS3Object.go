// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"io"
	"io/ioutil"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ReadS3Object returns a ByteReadCloser for an object in AWS S3.
// alg may be "bzip2", "gzip", "snappy", "zip", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func ReadS3Object(bucket string, key string, alg string, bufferSize int, s3Client *s3.S3) (ByteReadCloser, *Metadata, error) {

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := s3Client.GetObject(input)
	if err != nil {
		return &Reader{}, nil, errors.Wrap(err, "Error fetching data from S3")
	}

	metadata := NewMetadataFromS3(result)

	switch alg {
	case AlgorithmBzip2, AlgorithmGzip, AlgorithmSnappy, AlgorithmNone, "":
		r, closers, err := WrapReader(result.Body, []io.Closer{result.Body}, alg, bufferSize)
		if err != nil {
			return nil, metadata, errors.Wrapf(err, "error wrapping reader for object at uri s3://%s/%s", bucket, key)
		}
		return &Reader{Reader: r, Closers: closers}, metadata, nil
	case AlgorithmZip:
		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			return nil, metadata, errors.Wrapf(err, "error reading bytes from zip-compressed object at uri s3://%s/%s", bucket, key)
		}
		brc, err := ReadZipBytes(body)
		if err != nil {
			return nil, metadata, errors.Wrapf(err, "error creating reader for zip bytes for object at uri s3://%s/%s", bucket, key)
		}
		return brc, metadata, nil
	}
	return nil, metadata, &ErrUnknownAlgorithm{Algorithm: alg}

}

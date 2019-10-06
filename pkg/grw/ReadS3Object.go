// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-reader-writer/pkg/io"
)

type ReadS3ObjectInput struct {
	Bucket     string // the S3 bucket name
	Key        string // the key to the file in the bucket
	Alg        string // the compression algorithm
	Dict       []byte // the dictionary for the compression algorithm, if applicable
	BufferSize int    // the buffer size for the underlying reader
	S3Client   *s3.S3 // the S3 client
}

// ReadS3Object returns a ByteReadCloser for an object in AWS S3.
// alg may be "bzip2", "gzip", "snappy", "zip", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func ReadS3Object(input *ReadS3ObjectInput) (*Reader, *Metadata, error) {

	result, err := input.S3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(input.Bucket),
		Key:    aws.String(input.Key),
	})
	if err != nil {
		return &Reader{}, nil, errors.Wrap(err, "Error fetching data from S3")
	}

	metadata := NewMetadataFromS3(result)

	switch input.Alg {
	case AlgorithmBzip2, AlgorithmGzip, AlgorithmSnappy, AlgorithmNone, "":
		r, err := WrapReader(result.Body, input.Alg, input.Dict, input.BufferSize)
		if err != nil {
			return nil, metadata, errors.Wrapf(err, "error wrapping reader for object at uri s3://%s/%s", input.Bucket, input.Key)
		}
		return &Reader{Reader: r}, metadata, nil
	case AlgorithmZip:
		body, err := io.ReadAll(result.Body)
		if err != nil {
			return nil, metadata, errors.Wrapf(err, "error reading bytes from zip-compressed object at uri s3://%s/%s", input.Bucket, input.Key)
		}
		brc, err := ReadZipBytes(body)
		if err != nil {
			return nil, metadata, errors.Wrapf(err, "error creating reader for zip bytes for object at uri s3://%s/%s", input.Bucket, input.Key)
		}
		return brc, metadata, nil
	}
	return nil, metadata, &ErrUnknownAlgorithm{Algorithm: input.Alg}

}

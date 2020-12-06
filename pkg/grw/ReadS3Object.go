// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

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
		return &Reader{}, nil, fmt.Errorf("error fetching data from S3: %w", err)
	}

	metadata := NewMetadataFromS3(result)

	switch input.Alg {
	case AlgorithmBzip2, AlgorithmGzip, AlgorithmSnappy, AlgorithmNone, "":
		r, err := WrapReader(result.Body, input.Alg, input.Dict, input.BufferSize)
		if err != nil {
			return nil, metadata, fmt.Errorf("error wrapping reader for object at uri s3://%s/%s: %w", input.Bucket, input.Key, err)
		}
		return &Reader{Reader: r}, metadata, nil
	case AlgorithmZip:
		body, err := io.ReadAll(result.Body)
		if err != nil {
			return nil, metadata, fmt.Errorf("error reading bytes from zip-compressed object at uri s3://%s/%s: %w", input.Bucket, input.Key, err)
		}
		brc, err := ReadZipBytes(body)
		if err != nil {
			return nil, metadata, fmt.Errorf("error creating reader for zip bytes for object at uri s3://%s/%s: %w", input.Bucket, input.Key, err)
		}
		return brc, metadata, nil
	}
	return nil, metadata, &ErrUnknownAlgorithm{Algorithm: input.Alg}

}

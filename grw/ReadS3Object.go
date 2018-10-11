// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"archive/zip"
	"bufio"
	"bytes"
	"compress/bzip2"
	"compress/gzip"
	"io/ioutil"
)

import (
	"github.com/golang/snappy"
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
func ReadS3Object(bucket string, key string, alg string, cache bool, s3_client *s3.S3) (ByteReadCloser, *Metadata, error) {

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := s3_client.GetObject(input)
	if err != nil {
		return &Reader{}, nil, errors.Wrap(err, "Error fetching data from S3")
	}

	if alg == "gzip" {

		gr, err := gzip.NewReader(result.Body)
		if err != nil {
			return &Reader{}, nil, errors.Wrap(err, "Error creating gizp reader for AWS s3 object at s3://"+bucket+"/"+key+".")
		}

		if cache {
			return &Cache{
				Reader:  &Reader{Reader: bufio.NewReader(gr), Closer: gr},
				Content: &[]byte{},
			}, NewMetadataFromS3(result), nil
		}

		return &Reader{Reader: bufio.NewReader(gr), Closer: gr}, NewMetadataFromS3(result), nil

	}

	if alg == "bzip2" {

		br := bzip2.NewReader(result.Body)

		if cache {
			return &Cache{
				Reader:  &Reader{Reader: bufio.NewReader(br), Closer: result.Body},
				Content: &[]byte{},
			}, NewMetadataFromS3(result), nil
		}

		return &Reader{Reader: bufio.NewReader(br), Closer: result.Body}, NewMetadataFromS3(result), nil

	}

	if alg == "snappy" {

		sr := snappy.NewReader(bufio.NewReader(result.Body))

		if cache {
			return &Cache{
				Reader:  &Reader{Reader: bufio.NewReader(sr), Closer: result.Body},
				Content: &[]byte{},
			}, NewMetadataFromS3(result), nil
		}

		return &Reader{Reader: bufio.NewReader(sr), Closer: result.Body}, NewMetadataFromS3(result), nil
	}

	if alg == "zip" {

		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			return nil, nil, errors.Wrap(err, "error reading S3 object body")
		}

		zr, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
		if err != nil {
			return nil, nil, errors.Wrap(err, "Error opening zip file at \"s3://"+bucket+"/"+key+"\" for reading")
		}

		if len(zr.File) != 1 {
			return nil, nil, errors.New("error zip file has more than one internal file.")
		}

		zfr, err := zr.File[0].Open()
		if err != nil {
			return nil, nil, errors.Wrap(err, "error opening internal file for zip.")
		}

		if cache {
			return &Cache{
				Reader:  &Reader{Reader: bufio.NewReader(zfr), Closer: zfr},
				Content: &[]byte{},
			}, NewMetadataFromS3(result), nil
		}

		return &Reader{Reader: bufio.NewReader(zfr), Closer: zfr}, NewMetadataFromS3(result), nil

	}

	if cache {
		return &Cache{
			Reader:  &Reader{Reader: bufio.NewReader(result.Body), Closer: result.Body},
			Content: &[]byte{},
		}, NewMetadataFromS3(result), nil
	}

	return &Reader{Reader: bufio.NewReader(result.Body), Closer: result.Body}, NewMetadataFromS3(result), nil

}

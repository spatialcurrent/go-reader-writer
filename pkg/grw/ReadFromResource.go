// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

func ReadFromResource(uri string, alg string, bufferSize int, s3Client *s3.S3) (ByteReadCloser, *Metadata, error) {

	return NewBuilder().
		Uri(uri).
		Algorithm(alg).
		BufferSize(bufferSize).
		S3Client(s3Client).
		Open()

}

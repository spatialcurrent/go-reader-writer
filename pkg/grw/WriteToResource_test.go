// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pkgalg "github.com/spatialcurrent/go-reader-writer/pkg/alg"
)

func TestWriteToStdout(t *testing.T) {
	output, err := WriteToResource(&WriteToResourceInput{
		URI:        "file:///dev/stdout",
		Alg:        pkgalg.AlgorithmNone,
		Dict:       NoDict,
		BufferSize: 4096,
		S3Client:   nil,
	})
	assert.NoError(t, err)
	assert.NotNil(t, output.Writer)

	n, err := output.Writer.Write(BytesHelloWorld)
	assert.NoError(t, err)
	assert.Equal(t, len(BytesHelloWorld), n)
}

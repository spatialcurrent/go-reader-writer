// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"testing"

	pkgalg "github.com/spatialcurrent/go-reader-writer/pkg/alg"
	"github.com/stretchr/testify/assert"
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

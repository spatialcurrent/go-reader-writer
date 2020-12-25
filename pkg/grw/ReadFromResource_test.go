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
	"github.com/spatialcurrent/go-reader-writer/pkg/io"
	"github.com/stretchr/testify/assert"
)

func TestReadFromResourceDocTxt(t *testing.T) {
	output, err := ReadFromResource(&ReadFromResourceInput{
		URI:        "file://../../testdata/doc.txt",
		Alg:        pkgalg.AlgorithmNone,
		Dict:       NoDict,
		BufferSize: 4096,
		S3Client:   nil,
	})
	assert.NoError(t, err)
	assert.NotNil(t, output.Reader)
	assert.Nil(t, output.Metadata)

	got, err := io.ReadAllAndClose(output.Reader)
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromResourceDocTxtBz2(t *testing.T) {
	output, err := ReadFromResource(&ReadFromResourceInput{
		URI:        "file://../../testdata/doc.txt.bz2",
		Alg:        pkgalg.AlgorithmBzip2,
		Dict:       NoDict,
		BufferSize: 4096,
		S3Client:   nil,
	})
	assert.NoError(t, err)
	assert.NotNil(t, output.Reader)
	assert.Nil(t, output.Metadata)

	got, err := io.ReadAllAndClose(output.Reader)
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromResourceDocTxtGzip(t *testing.T) {
	output, err := ReadFromResource(&ReadFromResourceInput{
		URI:        "file://../../testdata/doc.txt.gz",
		Alg:        pkgalg.AlgorithmGzip,
		Dict:       NoDict,
		BufferSize: 4096,
		S3Client:   nil,
	})
	assert.NoError(t, err)
	assert.NotNil(t, output.Reader)
	assert.Nil(t, output.Metadata)

	got, err := io.ReadAllAndClose(output.Reader)
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromResourceDocTxtSnappy(t *testing.T) {
	output, err := ReadFromResource(&ReadFromResourceInput{
		URI:        "file://../../testdata/doc.txt.sz",
		Alg:        pkgalg.AlgorithmSnappy,
		Dict:       NoDict,
		BufferSize: 4096,
		S3Client:   nil,
	})
	assert.NoError(t, err)
	assert.NotNil(t, output.Reader)
	assert.Nil(t, output.Metadata)

	got, err := io.ReadAllAndClose(output.Reader)
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromResourceDocTxtZip(t *testing.T) {
	output, err := ReadFromResource(&ReadFromResourceInput{
		URI:        "file://../../testdata/doc.txt.zip",
		Alg:        pkgalg.AlgorithmZip,
		Dict:       NoDict,
		BufferSize: 4096,
		S3Client:   nil,
	})
	assert.NoError(t, err)
	assert.NotNil(t, output.Reader)
	assert.Nil(t, output.Metadata)

	got, err := io.ReadAllAndClose(output.Reader)
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

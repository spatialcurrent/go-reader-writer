// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFromResourceDocTxt(t *testing.T) {
	brc, metadata, err := ReadFromResource(&ReadFromResourceInput{
		Uri:        "../../testdata/doc.txt",
		Alg:        AlgorithmNone,
		Dict:       NoDict,
		BufferSize: 4096,
		S3Client:   nil,
	})
	assert.NoError(t, err)
	assert.NotNil(t, brc)
	assert.Nil(t, metadata)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromResourceDocTxtBz2(t *testing.T) {
	brc, metadata, err := ReadFromResource(&ReadFromResourceInput{
		Uri:        "../../testdata/doc.txt.bz2",
		Alg:        AlgorithmBzip2,
		Dict:       NoDict,
		BufferSize: 4096,
		S3Client:   nil,
	})
	assert.NoError(t, err)
	assert.NotNil(t, brc)
	assert.Nil(t, metadata)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromResourceDocTxtGzip(t *testing.T) {
	brc, metadata, err := ReadFromResource(&ReadFromResourceInput{
		Uri:        "../../testdata/doc.txt.gz",
		Alg:        AlgorithmGzip,
		Dict:       NoDict,
		BufferSize: 4096,
		S3Client:   nil,
	})
	assert.NoError(t, err)
	assert.NotNil(t, brc)
	assert.Nil(t, metadata)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromResourceDocTxtSnappy(t *testing.T) {
	brc, metadata, err := ReadFromResource(&ReadFromResourceInput{
		Uri:        "../../testdata/doc.txt.sz",
		Alg:        AlgorithmSnappy,
		Dict:       NoDict,
		BufferSize: 4096,
		S3Client:   nil,
	})
	assert.NoError(t, err)
	assert.NotNil(t, brc)
	assert.Nil(t, metadata)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromResourceDocTxtZip(t *testing.T) {
	brc, metadata, err := ReadFromResource(&ReadFromResourceInput{
		Uri:        "../../testdata/doc.txt.zip",
		Alg:        AlgorithmZip,
		Dict:       NoDict,
		BufferSize: 4096,
		S3Client:   nil,
	})
	assert.NoError(t, err)
	assert.NotNil(t, brc)
	assert.Nil(t, metadata)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

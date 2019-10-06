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

func TestReadFromFilePathDocTxt(t *testing.T) {
	brc, err := ReadFromFilePath(&ReadFromFilePathInput{Path: "../../testdata/doc.txt", Alg: AlgorithmNone, BufferSize: 4096})
	assert.NoError(t, err)
	assert.NotNil(t, brc)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromFilePathDocTxtBz2(t *testing.T) {
	brc, err := ReadFromFilePath(&ReadFromFilePathInput{Path: "../../testdata/doc.txt.bz2", Alg: AlgorithmBzip2, BufferSize: 4096})
	assert.NoError(t, err)
	assert.NotNil(t, brc)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromFilePathDocTxtGzip(t *testing.T) {
	brc, err := ReadFromFilePath(&ReadFromFilePathInput{Path: "../../testdata/doc.txt.gz", Alg: AlgorithmGzip, BufferSize: 4096})
	assert.NoError(t, err)
	assert.NotNil(t, brc)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromFilePathDocTxtSnappy(t *testing.T) {
	brc, err := ReadFromFilePath(&ReadFromFilePathInput{Path: "../../testdata/doc.txt.sz", Alg: AlgorithmSnappy, BufferSize: 4096})
	assert.NoError(t, err)
	assert.NotNil(t, brc)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromFilePathDocTxtZip(t *testing.T) {
	brc, err := ReadFromFilePath(&ReadFromFilePathInput{Path: "../../testdata/doc.txt.zip", Alg: AlgorithmZip, BufferSize: 4096})
	assert.NoError(t, err)
	assert.NotNil(t, brc)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

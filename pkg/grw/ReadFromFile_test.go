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

	"github.com/spatialcurrent/go-reader-writer/pkg/os"
)

func TestReadFromFileDocTxt(t *testing.T) {
	f, err := os.OpenFile("../../testdata/doc.txt")
	assert.NoError(t, err)

	brc, err := ReadFromFile(f, AlgorithmNone, NoDict, DefaultBufferSize)
	assert.NoError(t, err)
	assert.NotNil(t, brc)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromFileDocTxtBz2(t *testing.T) {
	f, err := os.OpenFile("../../testdata/doc.txt.bz2")
	assert.NoError(t, err)

	brc, err := ReadFromFile(f, AlgorithmBzip2, NoDict, DefaultBufferSize)
	assert.NoError(t, err)
	assert.NotNil(t, brc)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromFileDocTxtGzip(t *testing.T) {
	f, err := os.OpenFile("../../testdata/doc.txt.gz")
	assert.NoError(t, err)

	brc, err := ReadFromFile(f, AlgorithmGzip, NoDict, DefaultBufferSize)
	assert.NoError(t, err)
	assert.NotNil(t, brc)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromFileDocTxtSnappy(t *testing.T) {
	f, err := os.OpenFile("../../testdata/doc.txt.sz")
	assert.NoError(t, err)

	brc, err := ReadFromFile(f, AlgorithmSnappy, NoDict, DefaultBufferSize)
	assert.NoError(t, err)
	assert.NotNil(t, brc)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromFileDocTxtZip(t *testing.T) {
	f, err := os.OpenFile("../../testdata/doc.txt.zip")
	assert.NoError(t, err)

	brc, err := ReadFromFile(f, AlgorithmZip, NoDict, DefaultBufferSize)
	assert.NoError(t, err)
	assert.NotNil(t, brc)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

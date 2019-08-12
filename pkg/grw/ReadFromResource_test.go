// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestReadFromResourceDocTxt(t *testing.T) {
	brc, metadata, err := ReadFromResource("../../testdata/doc.txt", AlgorithmNone, 4096, nil)
	assert.NoError(t, err)
	assert.NotNil(t, brc)
	assert.Nil(t, metadata)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromResourceDocTxtBz2(t *testing.T) {
	brc, metadata, err := ReadFromResource("../../testdata/doc.txt.bz2", AlgorithmBzip2, 4096, nil)
	assert.NoError(t, err)
	assert.NotNil(t, brc)
	assert.Nil(t, metadata)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromResourceDocTxtGzip(t *testing.T) {
	brc, metadata, err := ReadFromResource("../../testdata/doc.txt.gz", AlgorithmGzip, 4096, nil)
	assert.NoError(t, err)
	assert.NotNil(t, brc)
	assert.Nil(t, metadata)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromResourceDocTxtSnappy(t *testing.T) {
	brc, metadata, err := ReadFromResource("../../testdata/doc.txt.sz", AlgorithmSnappy, 4096, nil)
	assert.NoError(t, err)
	assert.NotNil(t, brc)
	assert.Nil(t, metadata)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

func TestReadFromResourceDocTxtZip(t *testing.T) {
	brc, metadata, err := ReadFromResource("../../testdata/doc.txt.zip", AlgorithmZip, 4096, nil)
	assert.NoError(t, err)
	assert.NotNil(t, brc)
	assert.Nil(t, metadata)

	got, err := brc.ReadAllAndClose()
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)
}

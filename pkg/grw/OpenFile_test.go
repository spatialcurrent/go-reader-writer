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
	"github.com/stretchr/testify/require"
)

func TestOpenFileDocTxt(t *testing.T) {
	f, err := OpenFile("../../testdata/doc.txt")
	require.NoError(t, err)
	require.NotNil(t, f)
	defer f.Close()
	assert.Equal(t, "../../testdata/doc.txt", f.Name())
}

func TestOpenFileDocTxtBz2(t *testing.T) {
	f, err := OpenFile("../../testdata/doc.txt.bz2")
	require.NoError(t, err)
	require.NotNil(t, f)
	defer f.Close()
	assert.Equal(t, "../../testdata/doc.txt.bz2", f.Name())
}

func TestOpenFileDocTxtGzip(t *testing.T) {
	f, err := OpenFile("../../testdata/doc.txt.gz")
	require.NoError(t, err)
	require.NotNil(t, f)
	defer f.Close()
	assert.Equal(t, "../../testdata/doc.txt.gz", f.Name())
}

func TestOpenFileDocTxtSnappy(t *testing.T) {
	f, err := OpenFile("../../testdata/doc.txt.sz")
	require.NoError(t, err)
	require.NotNil(t, f)
	defer f.Close()
	assert.Equal(t, "../../testdata/doc.txt.sz", f.Name())
}

func TestOpenFileDocTxtZip(t *testing.T) {
	f, err := OpenFile("../../testdata/doc.txt.zip")
	require.NoError(t, err)
	require.NotNil(t, f)
	defer f.Close()
	assert.Equal(t, "../../testdata/doc.txt.zip", f.Name())
}

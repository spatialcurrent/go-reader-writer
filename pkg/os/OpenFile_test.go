// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package os

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func closeFile(t *testing.T, f *os.File) {
	err := f.Close()
	if err != nil {
		t.Error(fmt.Errorf("error closing file at path %q: %w", f.Name(), err).Error())
	}
}

func TestOpenFileDocTxt(t *testing.T) {
	f, err := OpenFile("../../testdata/doc.txt")
	require.NoError(t, err)
	require.NotNil(t, f)
	defer closeFile(t, f)
	assert.Equal(t, "../../testdata/doc.txt", f.Name())
}

func TestOpenFileDocTxtBz2(t *testing.T) {
	f, err := OpenFile("../../testdata/doc.txt.bz2")
	require.NoError(t, err)
	require.NotNil(t, f)
	defer closeFile(t, f)
	assert.Equal(t, "../../testdata/doc.txt.bz2", f.Name())
}

func TestOpenFileDocTxtGzip(t *testing.T) {
	f, err := OpenFile("../../testdata/doc.txt.gz")
	require.NoError(t, err)
	require.NotNil(t, f)
	defer closeFile(t, f)
	assert.Equal(t, "../../testdata/doc.txt.gz", f.Name())
}

func TestOpenFileDocTxtSnappy(t *testing.T) {
	f, err := OpenFile("../../testdata/doc.txt.sz")
	require.NoError(t, err)
	require.NotNil(t, f)
	defer closeFile(t, f)
	assert.Equal(t, "../../testdata/doc.txt.sz", f.Name())
}

func TestOpenFileDocTxtZip(t *testing.T) {
	f, err := OpenFile("../../testdata/doc.txt.zip")
	require.NoError(t, err)
	require.NotNil(t, f)
	defer closeFile(t, f)
	assert.Equal(t, "../../testdata/doc.txt.zip", f.Name())
}

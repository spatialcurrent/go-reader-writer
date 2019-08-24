// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestWriteGzipBytes(t *testing.T) {
	original := []byte("hello world")
	writer, buffer := WriteGzipBytes()
	n, err := writer.Write(original)
	assert.NoError(t, err)
	assert.Equal(t, 11, n)
	err = writer.Flush()
	assert.NoError(t, err)
	err = writer.Close()
	assert.NoError(t, err)
	data := buffer.Bytes()
	assert.Equal(
		t,
		[]byte{
			0x1f, 0x8b, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xca, 0x48, 0xcd,
			0xc9, 0xc9, 0x57, 0x28, 0xcf, 0x2f, 0xca, 0x49, 0x1, 0x4, 0x0, 0x0, 0xff,
			0xff, 0x85, 0x11, 0x4a, 0xd, 0xb, 0x0, 0x0, 0x0,
		},
		data,
	)
	r, err := gzip.NewReader(bytes.NewReader(data))
	assert.NoError(t, err)
	out, err := ioutil.ReadAll(r)
	assert.NoError(t, err)
	assert.Equal(t, original, out)
}

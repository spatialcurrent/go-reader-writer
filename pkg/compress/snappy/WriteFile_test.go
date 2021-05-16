// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package snappy

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteFile(t *testing.T) {
	_ = os.MkdirAll("temp", 0775)
	f, err := os.CreateTemp("temp", "*.sz")
	assert.NoError(t, err)
	defer removeFile(t, f.Name())

	w, err := WriteFile(f.Name(), 4096)
	assert.NoError(t, err)
	assert.NotNil(t, w)

	n, err := w.Write(BytesHelloWorld)
	assert.Equal(t, n, len(BytesHelloWorld))
	assert.NoError(t, err)

	err = w.Flush()
	assert.NoError(t, err)

	err = w.Close()
	assert.NoError(t, err)
}

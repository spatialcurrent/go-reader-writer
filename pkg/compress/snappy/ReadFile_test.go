// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package snappy

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	r, err := ReadFile("../../../testdata/doc.txt.sz", 4096)
	assert.NoError(t, err)
	assert.NotNil(t, r)

	got, err := ioutil.ReadAll(r)
	assert.NoError(t, err)
	assert.Equal(t, BytesHelloWorld, got)

	err = r.Close()
	assert.NoError(t, err)

	err = r.Close()
	assert.Error(t, err)
}

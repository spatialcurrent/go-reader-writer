// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package bufio

import (
	"bytes"
	"crypto/rand"
	"io/ioutil"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	f := func() bool {
		in := make([]byte, 10)
		_, err := rand.Read(in)
		assert.NoError(t, err)
		out, err := ioutil.ReadAll(NewReader(NewReader(ioutil.NopCloser(bytes.NewReader(in)))))
		assert.NoError(t, err)
		return bytes.Equal(in, out)
	}
	assert.NoError(t, quick.Check(f, nil))
}

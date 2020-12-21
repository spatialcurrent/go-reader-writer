// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gzip

import (
	"bytes"
	"crypto/rand"
	"io/ioutil"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
)

var (
	BytesHelloWorld = []byte("hello world")
)

func TestGzip(t *testing.T) {
	f := func() bool {

		//
		// Create random input
		//

		in := make([]byte, 8192)
		_, err := rand.Read(in)
		if !assert.NoError(t, err) {
			return false
		}

		//
		// Create Buffer
		//

		buf := new(bytes.Buffer)

		//
		// Create Writer
		//

		// wrap with bufio writer to test propagation.
		w := NewWriter(bufio.NewWriter(buf))

		// Write data to buffer
		_, err = w.Write(in)
		if !assert.NoError(t, err) {
			return false
		}

		// Flush all writers
		err = w.Flush()
		if !assert.NoError(t, err) {
			return false
		}

		// Close all writers (save gzip trailer)
		err = w.Close()
		if !assert.NoError(t, err) {
			return false
		}

		// wrap with bufio reader to test propagation.
		r, err := NewReader(bufio.NewReader(ioutil.NopCloser(buf)))
		if !assert.NoError(t, err) {
			return false
		}

		out, err := ioutil.ReadAll(r)
		if !assert.NoError(t, err) {
			return false
		}

		if !assert.Equal(t, in, out) {
			return false
		}

		return true
	}
	assert.NoError(t, quick.Check(f, nil))
}

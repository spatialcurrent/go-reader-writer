// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package flate

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
)

var (
	BytesHelloWorld = []byte("hello world")
)

func removeFile(t *testing.T, path string) {
	err := os.Remove(path)
	if err != nil {
		t.Error(fmt.Errorf("error removing file at path %q: %w", path, err).Error())
	}
}

func TestFlateMemory(t *testing.T) {
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
		w, err := NewWriter(bufio.NewWriter(buf))
		if !assert.NoError(t, err) {
			return false
		}

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

		// Close all writers
		err = w.Close()
		if !assert.NoError(t, err) {
			return false
		}

		// wrap with bufio reader to test propagation.
		r := NewReader(bufio.NewReader(io.NopCloser(buf)))

		out, err := io.ReadAll(r)
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

func TestFlateFile(t *testing.T) {
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
		// Create File
		//

		_ = os.MkdirAll("temp", 0775)

		f, err := os.CreateTemp("temp", "*.f")
		_ = f.Close()
		assert.NoError(t, err)

		defer removeFile(t, f.Name())

		//
		// Create Writer
		//

		// wrap with bufio writer to test propagation.
		w, err := WriteFile(f.Name(), 4096)
		if !assert.NoError(t, err) {
			return false
		}

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
		r, err := ReadFile(f.Name(), nil, 4096)
		if !assert.NoError(t, err) {
			return false
		}

		out, err := io.ReadAll(r)
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

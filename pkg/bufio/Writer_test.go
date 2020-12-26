// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package bufio

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriter(t *testing.T) {
	buf := new(bytes.Buffer)
	w := NewWriter(NewWriter(buf))
	fmt.Fprint(w, "hello world")
	assert.Equal(t, "", buf.String())
	err := w.Flush()
	assert.NoError(t, err)
	assert.Equal(t, "hello world", buf.String())
}

func TestWriterOpen(t *testing.T) {
	buf := new(bytes.Buffer)
	w := NewWriter(NewWriterClose(buf, false))
	fmt.Fprint(w, "hello world")
	assert.Equal(t, "", buf.String())
	err := w.Flush()
	assert.NoError(t, err)
	assert.Equal(t, "hello world", buf.String())
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package splitter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitURI(t *testing.T) {
	scheme, remainder := SplitUri("https://example.com")
	assert.Equal(t, "https", scheme)
	assert.Equal(t, "example.com", remainder)
}

func TestSplitURIPort(t *testing.T) {
	scheme, remainder := SplitUri("https://example.com:80")
	assert.Equal(t, "https", scheme)
	assert.Equal(t, "example.com:80", remainder)
}

func TestSplitURIPortPath(t *testing.T) {
	scheme, remainder := SplitUri("https://example.com:80/foo")
	assert.Equal(t, "https", scheme)
	assert.Equal(t, "example.com:80/foo", remainder)
}

func TestSplitURINoScheme(t *testing.T) {
	scheme, remainder := SplitUri("example.com")
	assert.Equal(t, "", scheme)
	assert.Equal(t, "example.com", remainder)
}

func TestSplitURINoSchemePath(t *testing.T) {
	scheme, remainder := SplitUri("example.com/foo/bar")
	assert.Equal(t, "", scheme)
	assert.Equal(t, "example.com/foo/bar", remainder)
}

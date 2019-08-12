// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestSplitUri(t *testing.T) {
	scheme, remainder := SplitUri("https://example.com")
	assert.Equal(t, "https", scheme)
	assert.Equal(t, "example.com", remainder)
}

func TestSplitUriPort(t *testing.T) {
	scheme, remainder := SplitUri("https://example.com:80")
	assert.Equal(t, "https", scheme)
	assert.Equal(t, "example.com:80", remainder)
}

func TestSplitAuthorityPortPath(t *testing.T) {
	scheme, remainder := SplitUri("https://example.com:80/foo")
	assert.Equal(t, "https", scheme)
	assert.Equal(t, "example.com:80/foo", remainder)
}

func TestSplitAuthorityNoScheme(t *testing.T) {
	scheme, remainder := SplitUri("example.com")
	assert.Equal(t, "", scheme)
	assert.Equal(t, "example.com", remainder)
}

func TestSplitAuthorityNoSchemePath(t *testing.T) {
	scheme, remainder := SplitUri("example.com/foo/bar")
	assert.Equal(t, "", scheme)
	assert.Equal(t, "example.com/foo/bar", remainder)
}

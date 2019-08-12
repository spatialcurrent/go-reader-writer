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

func TestSplitAuthorityHost(t *testing.T) {
	userinfo, host, port := SplitAuthority("example.com")
	assert.Equal(t, "", userinfo)
	assert.Equal(t, "example.com", host)
	assert.Equal(t, "", port)
}

func TestSplitAuthorityHostPort(t *testing.T) {
	userinfo, host, port := SplitAuthority("example.com:21")
	assert.Equal(t, "", userinfo)
	assert.Equal(t, "example.com", host)
	assert.Equal(t, "21", port)
}

func TestSplitAuthorityHostPortUser(t *testing.T) {
	userinfo, host, port := SplitAuthority("foo@example.com:21")
	assert.Equal(t, "foo", userinfo)
	assert.Equal(t, "example.com", host)
	assert.Equal(t, "21", port)
}

func TestSplitAuthorityHostUserPassword(t *testing.T) {
	userinfo, host, port := SplitAuthority("foo:bar@example.com")
	assert.Equal(t, "foo:bar", userinfo)
	assert.Equal(t, "example.com", host)
	assert.Equal(t, "", port)
}

func TestSplitAuthorityHostPortUserPassword(t *testing.T) {
	userinfo, host, port := SplitAuthority("foo:bar@example.com:21")
	assert.Equal(t, "foo:bar", userinfo)
	assert.Equal(t, "example.com", host)
	assert.Equal(t, "21", port)
}

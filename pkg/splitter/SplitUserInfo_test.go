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
	"github.com/stretchr/testify/require"
)

func TestSplitUserInfoUser(t *testing.T) {
	user, password, err := SplitUserInfo("foo")
	require.Nil(t, err)
	assert.Equal(t, "foo", user)
	assert.Equal(t, "", password)
}

func TestSplitUserInfoUserColon(t *testing.T) {
	user, password, err := SplitUserInfo("foo:")
	require.Nil(t, err)
	assert.Equal(t, "foo", user)
	assert.Equal(t, "", password)
}

func TestSplitUserInfoUserPassword(t *testing.T) {
	user, password, err := SplitUserInfo("foo:bar")
	require.Nil(t, err)
	assert.Equal(t, "foo", user)
	assert.Equal(t, "bar", password)
}

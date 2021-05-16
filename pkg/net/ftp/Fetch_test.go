// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package ftp

import (
	"errors"
	"io"
	"net/textproto"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetch530(t *testing.T) {
	if os.Getenv("TEST_ACC_FTP") == "1" {
		r, err := Fetch("ftp://ftp2.census.gov/robots.txt")
		require.Nil(t, r)
		require.NotNil(t, err)
		require.IsType(t, &textproto.Error{}, errors.Unwrap(err))
		assert.Equal(t, 530, errors.Unwrap(err).(*textproto.Error).Code)
	}
}

func TestFetchAnonymous(t *testing.T) {
	if os.Getenv("TEST_ACC_FTP") == "1" {
		r, err := Fetch("ftp://anonymous@ftp2.census.gov/robots.txt")
		require.NoError(t, err)
		require.NotNil(t, r)
		got, err := io.ReadAll(r)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.Len(t, got, 478)
		err = r.Close()
		assert.NoError(t, err)
	}
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"errors"
	"net/textproto"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadFromFTPFile530(t *testing.T) {
	if os.Getenv("TEST_ACC_FTP") == "1" {
		brc, err := ReadFTPFile("ftp://ftp.fbo.gov/FBOFeed20011227", AlgorithmNone, NoDict, 4096)
		require.Nil(t, brc)
		require.NotNil(t, err)
		require.IsType(t, &textproto.Error{}, errors.Unwrap(err))
		assert.Equal(t, 530, errors.Unwrap(err).(*textproto.Error).Code)
	}
}

func TestReadFromFTPFileAnonymous(t *testing.T) {
	if os.Getenv("TEST_ACC_FTP") == "1" {
		brc, err := ReadFTPFile("ftp://anonymous@ftp.fbo.gov/FBOFeed20011227", AlgorithmNone, NoDict, 4096)
		require.NoError(t, err)
		require.NotNil(t, brc)
		got, err := brc.ReadAllAndClose()
		assert.NoError(t, err)
		assert.NotNil(t, got)
	}
}

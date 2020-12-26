// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package sftp

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"golang.org/x/crypto/ssh"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spatialcurrent/go-reader-writer/pkg/net/ssh2"
)

func TestFetchPassword(t *testing.T) {
	if os.Getenv("TEST_ACC_SFTP") == "1" {
		r, err := Fetch(os.Getenv("TEST_ACC_SFTP_URI"))
		require.NoError(t, err)
		require.NotNil(t, r)
		got, err := ioutil.ReadAll(r)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, "test", strings.TrimSpace(string(got)))
		err = r.Close()
		assert.NoError(t, err)
		err = r.Close()
		assert.Error(t, err)
	}
}

func TestFetchKey(t *testing.T) {
	if os.Getenv("TEST_ACC_SFTP") == "1" {
		key, err := ssh2.LoadPrivateKey(os.Getenv("TEST_ACC_SFTP_KEY"))
		if err != nil {
			t.Fatal(err)
		}
		r, err := Fetch(os.Getenv("TEST_ACC_SFTP_URI"), func(config *ssh2.ClientConfig) error {
			config.Auth = []ssh.AuthMethod{
				ssh.PublicKeys(key),
			}
			return nil
		})
		require.NoError(t, err)
		require.NotNil(t, r)
		got, err := ioutil.ReadAll(r)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, "test", strings.TrimSpace(string(got)))
		err = r.Close()
		assert.NoError(t, err)
		err = r.Close()
		assert.Error(t, err)
	}
}

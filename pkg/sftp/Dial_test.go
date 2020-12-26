// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package sftp

import (
	"golang.org/x/crypto/ssh"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDialPassword(t *testing.T) {
	if os.Getenv("TEST_ACC_SFTP") == "1" {
		r, err := Dial(os.Getenv("TEST_ACC_SFTP_URI"))
		require.NoError(t, err)
		require.NotNil(t, r)
		err = r.Close()
		assert.NoError(t, err)
		err = r.Close()
		assert.Error(t, err)
	}
}

func TestDialKey(t *testing.T) {
	if os.Getenv("TEST_ACC_SFTP") == "1" {
		key, err := LoadPrivateKey(os.Getenv("TEST_ACC_SFTP_KEY"))
		if err != nil {
			t.Fatal(err)
		}
		r, err := Dial(os.Getenv("TEST_ACC_SFTP_URI"), func(config *ClientConfig) error {
			config.Auth = []ssh.AuthMethod{
				ssh.PublicKeys(key),
			}
			return nil
		})
		require.NoError(t, err)
		require.NotNil(t, r)
		err = r.Close()
		assert.NoError(t, err)
		err = r.Close()
		assert.Error(t, err)
	}
}

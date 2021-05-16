// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package http

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchHTTP(t *testing.T) {
	if os.Getenv("TEST_ACC_HTTP") == "1" {
		r, err := Fetch(os.Getenv("TEST_ACC_HTTP_URI"))
		require.NoError(t, err)
		require.NotNil(t, r)
		got, err := io.ReadAll(r)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		err = r.Close()
		assert.NoError(t, err)
	}
}

func TestFetchHTTPS(t *testing.T) {
	if os.Getenv("TEST_ACC_HTTPS") == "1" {
		r, err := Fetch(os.Getenv("TEST_ACC_HTTPS_URI"))
		require.NoError(t, err)
		require.NotNil(t, r)
		got, err := io.ReadAll(r)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		err = r.Close()
		assert.NoError(t, err)
	}
}

func TestFetchHTTPSWithOption(t *testing.T) {
	if os.Getenv("TEST_ACC_HTTPS") == "1" {
		r, err := Fetch(os.Getenv("TEST_ACC_HTTPS_URI"), func(client *Client) error {
			client.Transport = &http.Transport{
				TLSClientConfig: &tls.Config{
					KeyLogWriter: os.Stderr,
				},
			}
			return nil
		})
		require.NoError(t, err)
		require.NotNil(t, r)
		got, err := io.ReadAll(r)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		err = r.Close()
		assert.NoError(t, err)
	}
}

// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package sftp2

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/ssh"

	"github.com/spatialcurrent/go-reader-writer/pkg/net/ssh2"
)

// FetchWithKey parses the provided key,
// appends the key to the ClientOption options, and
// then calls Fetch with the uri and options.
//
func FetchWithKey(uri string, key []byte, options ...ssh2.ClientOption) (*Reader, error) {

	if len(key) == 0 {
		return nil, errors.New("missing private SSH key")
	}

	privateKey, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("error parsing private SSH key: %w", err)
	}

	options = append(options, func(config *ssh2.ClientConfig) error {
		config.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		}
		return nil
	})

	return Fetch(uri, options...)

}

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

// Fetch returns a Reader for a file at a given SFTP address.
// ReadFTPFile returns the Reader and error, if any.
//
// ReadFTPFile returns an error if the address cannot be dialed,
// the userinfo cannot be parsed,
// the user and password are invalid, or
// the file cannot be retrieved.
//
// If a private key is provided, the function authenticates with the server
// and encrypts the connection using the key.
//
func FetchWithKey(uri string, key []byte, options ...ssh2.ClientOption) (*Reader, error) {

	if key == nil || len(key) == 0 {
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

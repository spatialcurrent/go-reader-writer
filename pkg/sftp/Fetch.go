// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package sftp

import (
	"fmt"
	"strings"

	"github.com/pkg/sftp"

	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
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
func Fetch(uri string, options ...ClientOption) (*Reader, error) {

	sshClient, err := Dial(uri, options...)
	if err != nil {
		return nil, fmt.Errorf("error creating SSH client: %w", err)
	}

	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return nil, fmt.Errorf("error creating SFTP client: %w", err)
	}

	_, fullpath := splitter.SplitUri(uri)
	parts := strings.SplitN(fullpath, "/", 2)

	file, err := sftpClient.Open(parts[1])
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return NewReader(file, sftpClient, sshClient), nil

}

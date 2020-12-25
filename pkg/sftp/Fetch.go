// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package sftp

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/sftp"

	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

const (
	DefaultPort    = 22
	DefaultTimeout = 5 * time.Second
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

	_, fullpath := splitter.SplitUri(uri)

	parts := strings.SplitN(fullpath, "/", 2)
	authority, p := parts[0], parts[1]

	userinfo, host, port := splitter.SplitAuthority(authority)
	if len(port) == 0 {
		port = strconv.Itoa(DefaultPort)
	}

	sshClientConfig := &ClientConfig{
		ClientConfig: ssh.ClientConfig{
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         DefaultTimeout,
		},
	}

	if len(userinfo) > 0 {
		user, password, err := splitter.SplitUserInfo(userinfo)
		if err != nil {
			return nil, fmt.Errorf("error parsing user info %q: %w", userinfo, err)
		}
		if len(user) > 0 {
			sshClientConfig.User = user
		}
		if len(password) > 0 {
			sshClientConfig.Auth = []ssh.AuthMethod{
				ssh.Password(password),
			}
		}
	}

	for i, option := range options {
		err := option(sshClientConfig)
		if err != nil {
			return nil, fmt.Errorf("error running client option %d: %w", i, err)
		}
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), &sshClientConfig.ClientConfig)
	if err != nil {
		return nil, fmt.Errorf("error creating SSH client: %w", err)
	}

	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return nil, fmt.Errorf("error creating SFTP client: %w", err)
	}

	file, err := sftpClient.Open(p)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return NewReader(file, sftpClient, sshClient), nil

}

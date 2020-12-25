// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package sftp

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Reader implements the io.ReadCloser interface to enabling reading
// a remote file on a SFTP server and closing the underlying connections.
type Reader struct {
	sftpFile   *sftp.File
	sftpClient *sftp.Client
	sshClient  *ssh.Client
}

// Read implements the io.Reader interface.
func (r *Reader) Read(p []byte) (int, error) {
	return r.sftpFile.Read(p)
}

// Close closes the file reader, the SFTP connection, and the SSH connection.
func (r *Reader) Close() error {
	err := r.sftpFile.Close()
	if err != nil {
		_ = r.sftpClient.Close() // attempt to close the underlying SFTP connection
		_ = r.sshClient.Close()  // attempt to close the underlying SSH connection
		return fmt.Errorf("error closing SFTP file: %w", err)
	}
	err = r.sftpClient.Close()
	if err != nil {
		_ = r.sshClient.Close() // attempt to close the underlying SSH connection
		return fmt.Errorf("error closing SFTP client connection: %w", err)
	}
	err = r.sshClient.Close()
	if err != nil {
		return fmt.Errorf("error closing SSH client connection: %w", err)
	}
	return nil
}

// NewReader creates a new Reader for reading a file from a SFTP server.
func NewReader(file *sftp.File, sftpClient *sftp.Client, sshClient *ssh.Client) *Reader {
	return &Reader{sftpFile: file, sftpClient: sftpClient, sshClient: sshClient}
}

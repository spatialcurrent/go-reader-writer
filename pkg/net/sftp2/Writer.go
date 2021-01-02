// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package sftp2

import (
	"fmt"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Writer implements the io.WriteCloser interface to enabling writing
// a remote file on a SFTP server and closing the underlying connections.
type Writer struct {
	sftpFile   *sftp.File
	sftpClient *sftp.Client
	sshClient  *ssh.Client
}

// Write implements the io.Writer interface.
func (w *Writer) Write(b []byte) (int, error) {
	return w.sftpFile.Write(b)
}

// Close closes the file reader, the SFTP connection, and the SSH connection.
func (w *Writer) Close() error {
	err := w.sftpFile.Close()
	if err != nil {
		if w.sftpClient != nil {
			_ = w.sftpClient.Close() // attempt to close the underlying SFTP connection
		}
		if w.sshClient != nil {
			_ = w.sshClient.Close() // attempt to close the underlying SSH connection
		}
		return fmt.Errorf("error closing SFTP file: %w", err)
	}
	if w.sftpClient != nil {
		err = w.sftpClient.Close()
		if err != nil {
			_ = w.sshClient.Close() // attempt to close the underlying SSH connection
			return fmt.Errorf("error closing SFTP client connection: %w", err)
		}
	}
	if w.sshClient != nil {
		err = w.sshClient.Close()
		if err != nil {
			return fmt.Errorf("error closing SSH client connection: %w", err)
		}
	}
	return nil
}

// NewWriter creates a new WRiter for reading a file from a SFTP server.
func NewWriter(file *sftp.File, sftpClient *sftp.Client, sshClient *ssh.Client) *Writer {
	return &Writer{sftpFile: file, sftpClient: sftpClient, sshClient: sshClient}
}

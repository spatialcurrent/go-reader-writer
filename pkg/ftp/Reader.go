// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package ftp

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

// Reader implements the io.ReadCloser interface to enabling reading
// a remote file on a FTP server and closing the underlying connections.
type Reader struct {
	response *ftp.Response
	server   *ftp.ServerConn
}

// Read implements the io.Reader interface.
func (r *Reader) Read(p []byte) (int, error) {
	return r.response.Read(p)
}

// SetDeadline sets the deadlines associated with the connection.
func (r *Reader) SetDeadline(t time.Time) error {
	return r.response.SetDeadline(t)
}

// Close closes the reader and quits the underlying connection.
func (r *Reader) Close() error {
	err := r.response.Close()
	if err != nil {
		_ = r.server.Quit() // attempt to quit the underlying connection
		return fmt.Errorf("error closing data connection: %w", err)
	}
	err = r.server.Quit()
	if err != nil {
		return fmt.Errorf("error quitting underlying server connection: %w", err)
	}
	return nil
}

// NewReader creates a new Reader for reading from a FTP server.
func NewReader(response *ftp.Response, server *ftp.ServerConn) *Reader {
	return &Reader{response: response, server: server}
}

// NewReader creates a new Reader for reading from a FTP server with a set deadline.
func NewReaderWithDeadline(response *ftp.Response, server *ftp.ServerConn, deadline time.Time) (*Reader, error) {
	err := response.SetDeadline(deadline)
	if err != nil {
		return nil, fmt.Errorf("error setting deadline %q: %w", deadline, err)
	}
	return &Reader{response: response, server: server}, nil
}

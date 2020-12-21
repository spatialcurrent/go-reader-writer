// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"

	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

// ReadFTPFile returns a Reader for an object for given ftp, ftps, or sftp address, and buffer size.
// ReadFTPFile returns the Reader and error, if any.
//
// ReadFTPFile returns an error if the address cannot be dialed,
// the userinfo cannot be parsed,
// the user and password are invalid, or
// the file cannot be retrieved.
//
func ReadFTPFile(uri string, alg string, dict []byte, bufferSize int) (*Reader, error) {

	_, fullpath := splitter.SplitUri(uri)

	parts := strings.SplitN(fullpath, "/", 2)
	authority, p := parts[0], parts[1]

	userinfo, host, port := splitter.SplitAuthority(authority)
	if len(port) == 0 {
		port = "21"
	}

	conn, err := ftp.Dial(fmt.Sprintf("%s:%s", host, port), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, fmt.Errorf("error opening file from uri %q: %w", uri, err)
	}

	if len(userinfo) > 0 {
		user, password, err := splitter.SplitUserInfo(userinfo)
		if err != nil {
			return nil, fmt.Errorf("error parsing userinfo %q: %w", userinfo, err)
		}
		if len(user) > 0 {
			err = conn.Login(user, password)
			if err != nil {
				return nil, fmt.Errorf("error logging in with user %q: %w", user, err)
			}
		}
	}

	resp, err := conn.Retr(p)
	if err != nil {
		return nil, fmt.Errorf("error reading file from uri %q: %w", uri, err)
	}

	return resp, nil

}

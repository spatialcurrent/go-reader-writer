// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package ftp

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"

	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

const (
	DefaultPort    = 21
	DefaultTimeout = 5 * time.Second
)

// Fetch returns a Reader for an object for given FTP address.
// ReadFTPFile returns the Reader and error, if any.
//
// ReadFTPFile returns an error if the address cannot be dialed,
// the userinfo cannot be parsed,
// the user and password are invalid, or
// the file cannot be retrieved.
//
func Fetch(uri string) (*Reader, error) {

	_, fullpath := splitter.SplitURI(uri)

	parts := strings.SplitN(fullpath, "/", 2)
	authority, p := parts[0], parts[1]

	userinfo, host, port := splitter.SplitAuthority(authority)
	if len(port) == 0 {
		port = strconv.Itoa(DefaultPort)
	}

	conn, err := ftp.Dial(fmt.Sprintf("%s:%s", host, port), ftp.DialWithTimeout(DefaultTimeout))
	if err != nil {
		return nil, fmt.Errorf("error opening file from uri %q: %w", uri, err)
	}

	if len(userinfo) > 0 {
		user, password, errSplitUserInfo := splitter.SplitUserInfo(userinfo)
		if errSplitUserInfo != nil {
			return nil, fmt.Errorf("error parsing user info %q: %w", userinfo, errSplitUserInfo)
		}
		if len(user) > 0 {
			errLogin := conn.Login(user, password)
			if errLogin != nil {
				return nil, fmt.Errorf("error logging in with user %q: %w", user, errLogin)
			}
		}
	}

	resp, errRetr := conn.Retr(p)
	if errRetr != nil {
		return nil, fmt.Errorf("error reading file from uri %q: %w", uri, errRetr)
	}

	if resp == nil {
		return nil, fmt.Errorf("error reading file from uri %q: response is empty", uri)
	}

	return NewReader(resp, conn), nil

}

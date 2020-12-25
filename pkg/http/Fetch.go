// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package http

import (
	"fmt"
	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	DefaultPortHTTP  = 80
	DefaultPortHTTPS = 443
)

func splitPath(str string) (string, string) {
	if i := strings.Index(str, "/"); i != -1 {
		return str[0:i], str[i:]
	}
	return str, ""
}

// Fetch returns a Reader for an object for given HTTP address.
// Fetch returns the Reader and error, if any.
//
// Fetch returns an error if the address cannot be reached,
// the userinfo cannot be parsed,
// the user and password are invalid, or
// the file cannot be retrieved.
//
func Fetch(uri string, options ...ClientOption) (io.ReadCloser, error) {

	scheme, fullpath := splitter.SplitUri(uri)

	if scheme != "http" && scheme != "https" {
		return nil, fmt.Errorf("error reading file from uri %q: http.Fetch only supports schemes http and https", uri)
	}

	authority, p := splitPath(fullpath)

	userinfo, host, port := splitter.SplitAuthority(authority)
	if len(port) == 0 {
		if scheme == "https" {
			port = strconv.Itoa(DefaultPortHTTPS)
		} else {
			port = strconv.Itoa(DefaultPortHTTP)
		}
	}

	client := &Client{}

	for i, option := range options {
		err := option(client)
		if err != nil {
			return nil, fmt.Errorf("error running client option %d: %w", i, err)
		}
	}

	request, err := http.NewRequest("GET", fmt.Sprintf("%s://%s:%s%s", scheme, host, port, p), nil)

	if len(userinfo) > 0 {
		user, password, err := splitter.SplitUserInfo(userinfo)
		if err != nil {
			return nil, fmt.Errorf("error parsing user info %q: %w", userinfo, err)
		}
		if len(user) > 0 && len(password) > 0 {
			request.SetBasicAuth(user, password)
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error reading file from uri %q: %w", uri, err)
	}

	if response == nil {
		return nil, fmt.Errorf("error reading file from uri %q: response is empty", uri)
	}

	return response.Body, nil

}

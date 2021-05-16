// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package splitter

import (
	"fmt"
	"net/url"
	"strings"
)

// SplitUserInfo splits a user info string into a user and password.
// User and password are unescaped using PathUnescape found in "net/url".
// Returns the user, password, and error, if any.
// The input user info string could match one of the following.
//  - user
//  - user:
//  - user:password
func SplitUserInfo(userinfo string) (string, string, error) {
	if len(userinfo) == 0 {
		return "", "", nil
	}
	if !strings.Contains(userinfo, ":") {
		user, err := url.PathUnescape(userinfo)
		if err != nil {
			return "", "", fmt.Errorf("error parsing user %q: %w", userinfo, err)
		}
		return user, "", nil
	}
	parts := strings.SplitN(userinfo, ":", 2)
	user, err := url.PathUnescape(parts[0])
	if err != nil {
		return "", "", fmt.Errorf("error parsing user %q: %w", parts[0], err)
	}
	password, err := url.PathUnescape(parts[1])
	if err != nil {
		return user, "", fmt.Errorf("error parsing password: %w", err)
	}
	return user, password, nil
}

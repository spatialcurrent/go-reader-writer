// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package splitter

import (
	"strings"
)

// SplitAuthority splits an authority string into userinfo, host, and port.
func SplitAuthority(authority string) (string, string, string) {

	if i := strings.Index(authority, "@"); i != -1 {
		if j := strings.Index(authority[i+1:], ":"); j != -1 {
			return authority[0:i], authority[i+1:][0:j], authority[i+1:][j+1:]
		}
		return authority[0:i], authority[i+1:], ""
	}

	if i := strings.Index(authority, ":"); i != -1 {
		return "", authority[0:i], authority[i+1:]
	}

	return "", authority, ""
}

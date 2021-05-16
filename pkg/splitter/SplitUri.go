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

// SplitURI splits a uri string into a scheme and remainder.
// If no scheme is specified, then returns "" as the scheme and the original string.
func SplitURI(uri string) (string, string) {
	if i := strings.Index(uri, "://"); i != -1 {
		return uri[0:i], uri[i+3:]
	}
	return "", uri
}

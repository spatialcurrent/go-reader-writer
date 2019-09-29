// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package io

import (
	"io/ioutil"
)

// ReadAll reads from r until an error or EOF and returns the data it read. A successful call returns err == nil, not err == EOF. Because ReadAll is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.
// If the given reader is nil, returns ErrMissingReader error.
func ReadAll(r Reader) ([]byte, error) {
	if r == nil {
		return make([]byte, 0), ErrMissingReader
	}
	return ioutil.ReadAll(r)
}

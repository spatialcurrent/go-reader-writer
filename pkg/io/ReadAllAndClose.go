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

// ReadAllAndClose reads all the data from the reader and calls its close method (if it has one).
// ReadAllAndClose will attempt to close the reader even if there was an error during reading.
func ReadAllAndClose(r Reader) ([]byte, error) {
	if r == nil {
		return make([]byte, 0), ErrMissingReader
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		Close(r) // ignores error from close and returns error from ReadAll
		return b, err
	}
	err = Close(r)
	return b, err

}

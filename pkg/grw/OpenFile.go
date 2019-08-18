// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"os"
)

import (
	"github.com/pkg/errors"
)

// OpenFile returns wraps os.OpenFile
func OpenFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening file at %q for reading", path)
	}
	return f, nil
}

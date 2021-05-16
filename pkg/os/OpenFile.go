// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package os

import (
	"fmt"
	"os"
)

// OpenFile returns wraps os.OpenFile
func OpenFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("error opening file at %q for reading: %w", path, err)
	}
	return f, nil
}

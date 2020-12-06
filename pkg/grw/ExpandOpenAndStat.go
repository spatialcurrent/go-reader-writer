// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

// ExpandOpenAndStat expands a path, returns *os.File and os.FileInfo for a given file path and an error, if any.
//
func ExpandOpenAndStat(path string) (*os.File, os.FileInfo, error) {

	pathExpanded, err := homedir.Expand(path)
	if err != nil {
		return nil, nil, fmt.Errorf("error expanding path %q: %w", path, err)
	}

	file, err := os.OpenFile(filepath.Clean(pathExpanded), os.O_RDONLY, 0600)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file at path %q: %w", path, err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, nil, fmt.Errorf("error stating file at path %q: %w", path, err)
	}

	return file, fileInfo, nil
}

// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"os"
	"path/filepath"
)

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

// ExpandOpenAndStat expands a path, returns *os.File and os.FileInfo for a given file path and an error, if any.
//
func ExpandOpenAndStat(path string) (*os.File, os.FileInfo, error) {

	pathExpanded, err := homedir.Expand(path)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error expanding path "+path)
	}

	file, err := os.OpenFile(filepath.Clean(pathExpanded), os.O_RDONLY, 0600)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error opening file at path "+path)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, nil, errors.Wrap(err, "error stating file at path "+path)
	}

	return file, fileInfo, nil
}

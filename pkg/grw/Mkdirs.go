// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

// Mkdirs create the directory at the given path and all parent directories, if they do not exist.
// Mkdirs expands the home directory and resolves the path given.
func Mkdirs(p string) error {

	if len(p) == 0 {
		return ErrPathMissing
	}

	pathExpanded, err := homedir.Expand(p)
	if err != nil {
		return errors.Wrapf(err, "error expanding file path %q", p)
	}

	pathAbsolute, err := filepath.Abs(pathExpanded)
	if err != nil {
		return errors.Wrapf(err, "error resolving file path %q", pathAbsolute)
	}

	err = os.MkdirAll(pathAbsolute, 0750)
	if err != nil {
		return errors.Wrapf(err, "error creating parent directories for %q", p)
	}

	return nil
}

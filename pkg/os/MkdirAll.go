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
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

// MkdirAll creates a directory named path, along with any necessary parents, and returns nil, or else returns an error.
// Wraps the standard library os.MkdirAll function to add support for expanding the user's home directory.
// The permission bits perm (before umask) are used for all directories that MkdirAll creates.
// If path is already a directory, MkdirAll does nothing and returns nil.
// Mkdirs expands the home directory and resolves the path given.
// Flag is the permissions flag, e.g., 0770.
//
//  - https://godoc.org/pkg/os/#MkdirAll
func MkdirAll(p string, flag os.FileMode) error {

	if len(p) == 0 {
		return ErrPathMissing
	}

	pathExpanded, err := homedir.Expand(p)
	if err != nil {
		return fmt.Errorf("error expanding file path %q: %w", p, err)
	}

	pathAbsolute, err := filepath.Abs(pathExpanded)
	if err != nil {
		return fmt.Errorf("error resolving file path %q: %w", pathAbsolute, err)
	}

	err = os.MkdirAll(pathAbsolute, flag)
	if err != nil {
		return fmt.Errorf("error creating parent directories for %q: %w", p, err)
	}

	return nil
}

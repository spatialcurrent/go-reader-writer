// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package os

import (
	"os"

	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

// Info is a simple interface for returning file info.
type Info interface {
	Mode() os.FileMode
	Size() int64
}

// Stat stats the given resource.
// Returns a bool indicating whether the file exists, file info, and an error if any.
// If the underlying error was a "does not exist" error, then the error is supressed and returns false, nil, nil
// If the underlying error was any other type of error, then the existence value is not guaranteed.
// Do check the error returned and do not ignore the error with "exists, _, _ := Stat(/path/tofile)".
func Stat(uri string) (bool, Info, error) {

	if uri == "stdin" {
		info, err := os.Stdin.Stat()
		return info.Mode()&os.ModeNamedPipe != 0, info, err
	} else if uri == "stdout" {
		info, err := os.Stdout.Stat()
		return true, info, err
	} else if uri == "stderr" {
		info, err := os.Stderr.Stat()
		return true, info, err
	}

	scheme, path := splitter.SplitUri(uri)

	switch scheme {
	case "file", "none", "":
		info, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				return false, nil, nil
			}
			return false, nil, err
		}
		return true, info, err
	}

	return false, nil, errors.Errorf("could not stat path %q: unsupported scheme %q", path, scheme)
}

// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package sftp2

import (
	"os"

	"github.com/pkg/sftp"

	"github.com/spatialcurrent/go-reader-writer/pkg/stat"
)

// Stat stats the given resource.
// Returns a bool indicating whether the file exists, file info, and an error if any.
// If the underlying error was a "does not exist" error, then the error is supressed and returns false, nil, nil
// If the underlying error was any other type of error, then the existence value is not guaranteed.
// Do check the error returned and do not ignore the error with "exists, _, _ := Stat(/path/tofile)".
func Stat(client *sftp.Client, path string) (bool, stat.Info, error) {

	info, err := client.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil, nil
		}
		return false, nil, err
	}
	return true, stat.NewFileInfo(info), err
}

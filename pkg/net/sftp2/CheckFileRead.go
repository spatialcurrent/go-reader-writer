// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package sftp2

import (
	"fmt"

	"github.com/pkg/sftp"
)

func CheckFileRead(client *sftp.Client, path string) error {
	exists, fileInfo, err := Stat(client, path)
	if err != nil {
		return fmt.Errorf("error stating file at path %q: %w", path, err)
	}
	if !exists {
		return fmt.Errorf("file at path %q does not exist", path)
	}
	if !(fileInfo.IsRegular() || fileInfo.IsNamedPipe()) {
		return fmt.Errorf("file at path %q is neither a regular file or named pipe: %w", path, err)
	}
	return nil
}

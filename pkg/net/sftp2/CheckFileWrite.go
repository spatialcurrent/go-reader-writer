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

func CheckFileWrite(client *sftp.Client, path string, appendToFile bool, overwrite bool) error {
	if (!overwrite) && (!appendToFile) {
		exists, fileInfo, err := Stat(client, path)
		if err != nil {
			return fmt.Errorf("error stating file at path %q: %w", path, err)
		}
		if exists && (!fileInfo.IsDevice()) && (!fileInfo.IsNamedPipe()) {
			return fmt.Errorf("file already exists at path %q and neither append or overwrite is set", path)
		}
	}
	return nil
}

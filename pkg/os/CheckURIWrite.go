// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package os

import (
	"fmt"
)

func CheckURIWrite(uri string, appendToFile bool, overwrite bool) error {
	if (!overwrite) && (!appendToFile) {
		exists, fileInfo, err := Stat(uri)
		if err != nil {
			return fmt.Errorf("error stating file at uri %q: %w", uri, err)
		}
		if exists && (!fileInfo.IsDevice()) && (!fileInfo.IsNamedPipe()) {
			return fmt.Errorf("file already exists at uri %q and neither append or overwrite is set", uri)
		}
	}
	return nil
}

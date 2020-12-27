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

func CheckURIRead(uri string) error {
	exists, fileInfo, err := Stat(uri)
	if err != nil {
		return fmt.Errorf("error stating file at uri %q: %w", uri, err)
	}
	if !exists {
		return fmt.Errorf("file at uri %q does not exist", uri)
	}
	if !(fileInfo.IsRegular() || fileInfo.IsNamedPipe()) {
		return fmt.Errorf("file at uri %q is neither a regular file or named pipe: %w", uri, err)
	}
	return nil
}

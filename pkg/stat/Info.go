// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package stat

import (
	"os"
)

// Info is a simple interface for returning file info.
type Info interface {
	IsRegular() bool
	IsDevice() bool
	IsCharacterDevice() bool
	IsNamedPipe() bool
	Mode() os.FileMode
	Perm() os.FileMode
	Size() int64
}

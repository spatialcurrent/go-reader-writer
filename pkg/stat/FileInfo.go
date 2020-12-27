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

// FileInfo wraps os.FileInfo and adds some additional functions.
type FileInfo struct {
	FileInfo os.FileInfo
}

// IsDir returns true if the file is a directory.
func (f *FileInfo) IsDir() bool {
	return f.FileInfo.Mode().IsDir()
}

// IsRegular returns true if the file is a regular file.
func (f *FileInfo) IsRegular() bool {
	return f.FileInfo.Mode().IsRegular()
}

// IsDevice returns true if the file is a device.
func (f *FileInfo) IsDevice() bool {
	return f.FileInfo.Mode()&os.ModeDevice != 0
}

func (f *FileInfo) IsCharacterDevice() bool {
	return f.FileInfo.Mode()&os.ModeCharDevice != 0
}

// IsNamedPipe returns true if the file is a named pipe.
// Returns true is a device is available as a named pipe within the context of the current process.
// Only if stdin has data to be read, will it return true (including as a path /dev/stdin).
// Stdout will either return as a named pipe or device depending on the context.
func (f *FileInfo) IsNamedPipe() bool {
	return f.FileInfo.Mode()&os.ModeNamedPipe != 0
}

// Mode returns the file mode.
func (f *FileInfo) Mode() os.FileMode {
	return f.FileInfo.Mode()
}

// Perm returns the file mode permissions bits.
func (f *FileInfo) Perm() os.FileMode {
	return f.FileInfo.Mode().Perm()
}

// Size retuns the file size as int64.
func (f *FileInfo) Size() int64 {
	return f.FileInfo.Size()
}

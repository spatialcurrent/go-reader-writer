// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package os

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// OpenDevice returns a pointer to the device indicated by name.
// Matches the following names as case insensitive:
//  - stdout, /dev/stdout => os.Stdout
//  - stderr, /dev/stderr => os.Stderr
//  - stdin, /dev/stdin => os.Stdin
//  - null, /dev/null => ioutil.Discard
func OpenDevice(name string) io.Writer {
	switch strings.ToLower(name) {
	case "stdout", "/dev/stdout":
		return os.Stdout
	case "stderr", "/dev/stderr":
		return os.Stderr
	case "stdin", "/dev/stdin":
		return os.Stdin
	case "null", "/dev/null":
		return ioutil.Discard
	}
	return nil
}

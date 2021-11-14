// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package os provides a few functions used throughout the grw package for interfacing with the local file system.
package os

import (
	"errors"
	"os"
)

const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	//lint:ignore ST1003 keep identical to standard library
	O_RDONLY int = os.O_RDONLY // open the file read-only.
	//lint:ignore ST1003 keep identical to standard library
	O_WRONLY int = os.O_WRONLY // open the file write-only.
	//lint:ignore ST1003 keep identical to standard library
	O_RDWR int = os.O_RDWR // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	//lint:ignore ST1003 keep identical to standard library
	O_APPEND int = os.O_APPEND // append data to the file when writing.
	//lint:ignore ST1003 keep identical to standard library
	O_CREATE int = os.O_CREATE // create a new file if none exists.
	//lint:ignore ST1003 keep identical to standard library
	O_EXCL int = os.O_EXCL // used with O_CREATE, file must not exist.
	//lint:ignore ST1003 keep identical to standard library
	O_SYNC int = os.O_SYNC // open for synchronous I/O.
	//lint:ignore ST1003 keep identical to standard library
	O_TRUNC int = os.O_TRUNC // truncate regular writable file when opened.
)

var (
	ErrPathMissing = errors.New("path missing")
)

var (
	Stdin  = os.Stdin
	Stdout = os.Stdout
	Stderr = os.Stderr
)

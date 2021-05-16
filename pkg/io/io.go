// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package io supplements the interfaces provided by the standard library io package and provides functions for using those interfaces.
// The interfaces and functions in this package are used throughout the grw package.
package io

import (
	"errors"
	"io"
)

var (
	ErrMissingWriter   = errors.New("missing writer")
	ErrMissingReader   = errors.New("missing reader")
	ErrMissingResource = errors.New("missing resource")
)

var (
	EOF = io.EOF
)

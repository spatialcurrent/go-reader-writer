// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"

	"github.com/spatialcurrent/go-reader-writer/pkg/compress/flate"
)

// FlateBytes returns a reader for reading flate bytes from an input array.
// Wraps the "compress/flate" package.
//
//  - https://golang.org/pkg/compress/flate/
//
func ReadFlateBytes(b []byte, dict []byte) *Reader {
	return &Reader{Reader: bufio.NewReader(flate.ReadBytes(b, dict))}
}

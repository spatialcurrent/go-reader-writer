// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bytes"

	"github.com/spatialcurrent/go-reader-writer/pkg/bufio"
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/gzip"
)

// WriteGzipBytes returns a reader for reading the bytes from an input array, and an error if any.
func WriteGzipBytes() (*Writer, *bytes.Buffer) {
	buf := new(bytes.Buffer)
	return NewWriter(bufio.NewWriter(gzip.NewWriter(buf))), buf
}

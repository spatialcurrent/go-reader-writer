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
)

// WriteSnappyBytes returns a writer and buffer for writing uncompressed bytes.
func WriteMemoryBytes() (*Writer, *bytes.Buffer) {
	buf := new(bytes.Buffer)
	return NewWriter(bufio.NewWriter(buf)), buf
}

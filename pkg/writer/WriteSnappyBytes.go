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
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/snappy"
)

// WriteSnappyBytes returns a writer and buffer for writing snappy compressed bytes.
func WriteSnappyBytes() (*Writer, *bytes.Buffer) {
	buf := new(bytes.Buffer)
	return NewWriter(bufio.NewWriter(snappy.NewBufferedWriter(buf))), buf
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bytes"
	"fmt"
)

// WriteZlibBytes returns a reader for reading the bytes from an input slice, and an error if any.
func WriteZlibBytes(dict []byte) (*Writer, *bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	w, err := WrapWriter(buf, AlgorithmZlib, dict)
	if err != nil {
		return nil, nil, fmt.Errorf("error wrapping writer: %w", err)
	}
	return w, buf, nil
}

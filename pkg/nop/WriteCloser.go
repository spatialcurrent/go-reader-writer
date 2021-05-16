// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package nop

import (
	"io"
)

type WriteCloser interface {
	io.WriteCloser
}

type writeCloser struct {
	io.Writer
}

func (c *writeCloser) Close() error {
	return nil
}

func NewWriteCloser(r io.Writer) WriteCloser {
	return &writeCloser{r}
}

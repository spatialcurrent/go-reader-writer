// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package nop

import (
	"io"
)

type ReadCloser interface {
	io.ReadCloser
}

type readCloser struct {
	io.Reader
}

func (c *readCloser) Close() error {
	return nil
}

func NewReadCloser(r io.Reader) ReadCloser {
	return &readCloser{r}
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"io"
)

import (
	"github.com/pkg/errors"
)

// Closer is a helper struct for closing a sequential list of closers associated with a resource.
// This is used for flushing the footer for a given file compression alogrithm before closing the underlying file.
type Closer struct {
	Closers []io.Closer // underlying list of io.Closer
}

// Close closes all the underlying io.Closer sequentially.
func (c *Closer) Close() error {
	if c.Closers != nil {
		for i, x := range c.Closers {
			err := x.Close()
			if err != nil {
				return errors.Wrapf(err, "error closing underlying closer %d", i)
			}
		}
	}
	return nil
}

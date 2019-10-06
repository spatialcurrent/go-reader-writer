// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package io

import (
	"github.com/pkg/errors"
)

// Flush flushes the given writer if it has a Flush method.
// If the given writer does not have a Flush method, then it simply returns nil.
// If the given writer is nil, then returns the ErrMissingWriter error.
func Flush(w interface{}) error {
	if w == nil {
		return ErrMissingWriter
	}
	if f, ok := w.(Flusher); ok {
		err := f.Flush()
		if err != nil {
			return errors.Wrapf(err, "error flushing writer")
		}
	}
	return nil
}

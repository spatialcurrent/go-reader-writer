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

// Close closes the given resource if it has a Close method.
// If the given resource does not have a Close method, then it simply returns nil.
// If the given resource is nil, then returns the ErrMissingResource error.
func Close(i interface{}) error {
	if i == nil {
		return ErrMissingResource
	}
	if f, ok := i.(Closer); ok {
		err := f.Close()
		if err != nil {
			return errors.Wrapf(err, "error closing resource")
		}
	}
	return nil
}

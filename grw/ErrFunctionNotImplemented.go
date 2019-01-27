// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

type ErrFunctionNotImplemented struct {
	Function string
	Object   string
}

func (e *ErrFunctionNotImplemented) Error() string {
	return e.Function + " is not implemented by " + e.Object
}

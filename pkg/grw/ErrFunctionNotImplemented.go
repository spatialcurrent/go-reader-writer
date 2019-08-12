// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
)

type ErrFunctionNotImplemented struct {
	Function string
	Object   string
}

func (e *ErrFunctionNotImplemented) Error() string {
	return fmt.Sprintf("%s is not implemented by %s", e.Function, e.Object)
}

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

type ErrReaderNotImplemented struct {
	Algorithm string
}

func (e *ErrReaderNotImplemented) Error() string {
	return fmt.Sprintf("%s is not implemented for reading", e.Algorithm)
}

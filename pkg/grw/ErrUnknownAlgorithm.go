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

type ErrUnknownAlgorithm struct {
	Algorithm string
}

func (e *ErrUnknownAlgorithm) Error() string {
	return fmt.Sprintf("algorithm %q is not known", e.Algorithm)
}

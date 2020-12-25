// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package schemes

import (
	"fmt"
)

type ErrUnknownScheme struct {
	Scheme string
}

func (e *ErrUnknownScheme) Error() string {
	return fmt.Sprintf("scheme %q is not known", e.Scheme)
}

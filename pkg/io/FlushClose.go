// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package io

import (
	"fmt"
)

// FlushClose flushes and then closes the writer.
func FlushClose(w Writer) error {
	err := Flush(w)
	if err != nil {
		return fmt.Errorf("error flushing writer: %w", err)
	}
	err = Close(w)
	if err != nil {
		return fmt.Errorf("error closing writer: %w", err)
	}
	return nil
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"

	"github.com/spatialcurrent/go-reader-writer/pkg/io"
)

type CloseWritersInput struct {
	Writers map[string]io.Writer
	Flush   bool
}

// CloseWriters closes a map of writers.
// If flush is true, will flush each writer before closing it.
func CloseWriters(input *CloseWritersInput) error {

	for _, w := range input.Writers {

		if input.Flush {
			err := io.Flush(w)
			if err != nil {
				return fmt.Errorf("error flushing writer: %w", err)
			}
		}

		err := io.Close(w)
		if err != nil {
			return fmt.Errorf("error closing writer: %w", err)
		}

	}

	return nil
}

// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package bytes

import (
	"fmt"

	"github.com/spatialcurrent/go-reader-writer/pkg/compress/gzip"
)

// ReadGzipBytes returns a reader for reading gzip bytes from an input array.
// Wraps the "compress/gzip" package.
//
//  - https://pkg.go.dev/pkg/compress/gzip/
//
func ReadGzipBytes(b []byte, multistream bool) (gzip.ReadResetCloser, error) {
	gr, err := gzip.ReadBytes(b, multistream)
	if err != nil {
		return nil, fmt.Errorf("error creating gzip reader for memory block: %w", err)
	}
	return gr, nil
}

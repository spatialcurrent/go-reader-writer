// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package bytes

import (
	"github.com/spatialcurrent/go-reader-writer/pkg/compress/snappy"
)

// ReadSnappyBytes returns a reader for an input of snappy-compressed bytes, and an error if any.
//
//  - https://pkg.go.dev//github.com/golang/snappy
//
func ReadSnappyBytes(b []byte) snappy.ReadResetter {
	return snappy.ReadBytes(b)
}

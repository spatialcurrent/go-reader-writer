// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package bytes

import (
	"bytes"
)

// ReadPlainBytes returns a reader for reading the bytes from an input array, and an error if any.
func ReadPlainBytes(b []byte) ByteReadScanner {
	return bytes.NewReader(b)
}

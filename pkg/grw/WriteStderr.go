// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"os"

	"github.com/pkg/errors"
)

// WriteStderr returns a ByteWriteCloser for stderr with the given compression.
// alg may be "bzip2", "gzip", "snappy", "zip", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/flate/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//  - https://golang.org/pkg/archive/zip/
//
func WriteStderr(alg string, dict []byte) (*Writer, error) {
	w, err := WrapWriter(os.Stderr, alg, dict)
	if err != nil {
		return nil, errors.Wrap(err, "error wrapping stderr")
	}
	return w, nil
}

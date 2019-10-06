// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package io

// ReadString returns a string of all the bytes up to and including the first occurrence of the given delimiter and an error, if any.
// If the given reader is nil, returns ErrMissingReader error.
func ReadString(r ByteReader, delim byte) (string, error) {
	if r == nil {
		return "", ErrMissingReader
	}
	b, err := r.ReadBytes(delim)
	if err != nil {
		return "", err
	}
	return string(b), err
}

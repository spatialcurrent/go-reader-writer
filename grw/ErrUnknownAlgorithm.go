// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

type ErrUnknownAlgorithm struct {
	Algorithm string
}

func (e *ErrUnknownAlgorithm) Error() string {
	return e.Algorithm + " is not known"
}

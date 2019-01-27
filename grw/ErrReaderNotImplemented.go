// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

type ErrReaderNotImplemented struct {
	Algorithm string
}

func (e *ErrReaderNotImplemented) Error() string {
	return e.Algorithm + " is not implemented for reading"
}

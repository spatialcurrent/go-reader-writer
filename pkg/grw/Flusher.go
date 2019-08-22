// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

// Flusher is an interface that provides a function to flush a buffer to an underlying writer.
type Flusher interface {
	Flush() error
}

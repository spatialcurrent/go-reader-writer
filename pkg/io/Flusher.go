// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package io

// Flusher is a simple interface that provides the common Flush function used by buffered writers.
type Flusher interface {
	Flush() error
}

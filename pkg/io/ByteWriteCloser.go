// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package io

// ByteWriteCloser is an interface that supports writing bytes.
// ByteWriteCloser extends ByteWriter, Flusher, and io.Closer interfaces.
// ByteWriterCloser provides a Lock and Unlock function for interacting with a mutex to lock the writer for exclusive use and prevent concurrency errors.
type ByteWriteCloser interface {
	ByteWriter
	Flusher
	Closer
	Lock()            // lock the writer for exclusive use
	Unlock()          // unlock the writer
	FlushSafe() error // lock the writer before flushing (to prvent another routine from writing concurrently).
	CloseSafe() error // lock the writer before closing (to prevent another routine from writing concurrently)
	WriteString(s string) (n int, err error)
	WriteLine(s string) (n int, err error)
	WriteLineSafe(s string) (n int, err error)
	WriteError(e error) (n int, err error)
	WriteErrorSafe(e error) (n int, err error)
}

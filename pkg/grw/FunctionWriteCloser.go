// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

type FunctionWriteCloser struct {
	Writer func(p []byte) (n int, err error)
	Closer func() error
}

func (w *FunctionWriteCloser) Write(p []byte) (n int, err error) {
	return w.Writer(p)
}

func (w *FunctionWriteCloser) Close() error {
	return w.Closer()
}

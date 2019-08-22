// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grwjs

import (
	"github.com/spatialcurrent/go-reader-writer/pkg/grw"
)

type Reader struct {
	Reader grw.ByteReadCloser
}

func (r *Reader) Read(p []byte) map[string]interface{} {
	n, err := r.Reader.Read(p)
	return map[string]interface{}{
		"n":   n,
		"err": err,
	}
}

func (r *Reader) ReadByte() map[string]interface{} {
	b, err := r.Reader.ReadByte()
	return map[string]interface{}{
		"b":   b,
		"err": err,
	}
}

func (r *Reader) ReadAt(p []byte, off int64) map[string]interface{} {
	n, err := r.Reader.ReadAt(p, off)
	return map[string]interface{}{
		"n":   n,
		"err": err,
	}
}

func (r *Reader) ReadAll() map[string]interface{} {
	data, err := r.Reader.ReadAll()
	return map[string]interface{}{
		"data": data,
		"err":  err,
	}
}

func (r *Reader) Close() map[string]interface{} {
	err := r.Reader.Close()
	return map[string]interface{}{
		"err": err,
	}
}

func (r *Reader) ReadAllAndClose() map[string]interface{} {
	data, err := r.Reader.ReadAllAndClose()
	return map[string]interface{}{
		"data": data,
		"err":  err,
	}
}

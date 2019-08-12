// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// grw.so creates a shared library of Go that can be called by C, C++, or Python
//
//  - https://godoc.org/cmd/cgo
//  - https://blog.golang.org/c-go-cgo
//
package main

import (
	"C"
	"github.com/pkg/errors"
	"github.com/spatialcurrent/go-reader-writer/pkg/grw"
	"unsafe"
)

func main() {}

//export ReadString
func ReadString(uri *C.char, alg *C.char, str **C.char) *C.char {

	r, _, err := grw.ReadFromResource(C.GoString(uri), C.GoString(alg), 4096, nil)
	if err != nil {
		return C.CString(errors.Wrap(err, "error opening resource from uri "+C.GoString(uri)).Error())
	}

	b, err := r.ReadAll()
	if err != nil {
		return C.CString(errors.Wrap(err, "error reading from resource").Error())
	}

	*str = C.CString(string(b))

	return nil
}

//export ReadBytes
func ReadBytes(uri *C.char, alg *C.char, bytes *unsafe.Pointer, length *C.int) *C.char {

	r, _, err := grw.ReadFromResource(C.GoString(uri), C.GoString(alg), 4096, nil)
	if err != nil {
		return C.CString(errors.Wrap(err, "error opening resource from uri "+C.GoString(uri)).Error())
	}

	b, err := r.ReadAll()
	if err != nil {
		return C.CString(errors.Wrap(err, "error reading from resource").Error())
	}

	*bytes = unsafe.Pointer(&b[0])
	*length = C.int(len(b))

	return nil
}

//export WriteString
func WriteString(uri *C.char, alg *C.char, appendFlag C.int, contents *C.char) *C.char {

	w, err := grw.WriteToResource(C.GoString(uri), C.GoString(alg), appendFlag > 0, nil)
	if err != nil {
		return C.CString(errors.Wrap(err, "error opening resource from uri "+C.GoString(uri)).Error())
	}

	_, err = w.WriteString(C.GoString(contents))
	if err != nil {
		return C.CString(errors.Wrap(err, "error writing to resource").Error())
	}

	return nil
}

//export WriteBytes
func WriteBytes(uri *C.char, alg *C.char, appendFlag C.int, bytes unsafe.Pointer, length C.int) *C.char {

	w, err := grw.WriteToResource(C.GoString(uri), C.GoString(alg), appendFlag > 0, nil)
	if err != nil {
		return C.CString(errors.Wrap(err, "error opening resource from uri "+C.GoString(uri)).Error())
	}

	_, err = w.Write(C.GoBytes(bytes, length))
	if err != nil {
		return C.CString(errors.Wrap(err, "error writing to resource").Error())
	}

	return nil
}

//export Version
func Version() *C.char {
	return C.CString(grw.Version)
}

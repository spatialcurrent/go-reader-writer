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
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spatialcurrent/go-reader-writer/pkg/grw"
	"github.com/spatialcurrent/go-reader-writer/pkg/schemes"
	//"unsafe"
)

var gitBranch string
var gitCommit string

func main() {}

//export Version
func Version() *C.char {
	var b strings.Builder
	if len(gitBranch) > 0 {
		b.WriteString(fmt.Sprintf("Branch: %q\n", gitBranch))
	}
	if len(gitCommit) > 0 {
		b.WriteString(fmt.Sprintf("Commit: %q\n", gitCommit))
	}
	return C.CString(b.String())
}

//export Algorithms
func Algorithms() *C.char {
	return C.CString(strings.Join(grw.Algorithms, ","))
}

//export Schemes
func Schemes() *C.char {
	return C.CString(strings.Join(schemes.Schemes, ","))
}

//export ReadString
func ReadString(uri *C.char, alg *C.char, str **C.char) *C.char {

	readFromResourceOutput, err := grw.ReadFromResource(&grw.ReadFromResourceInput{
		URI: C.GoString(uri),
		Alg: C.GoString(alg),
	})
	if err != nil {
		return C.CString(fmt.Errorf("error opening resource at uri %q: %w", C.GoString(uri), err).Error())
	}

	b, err := ioutil.ReadAll(readFromResourceOutput.Reader)
	if err != nil {
		return C.CString(fmt.Errorf("error reading resource at uri %q: %w", C.GoString(uri), err).Error())
	}

	err = readFromResourceOutput.Reader.Close()
	if err != nil {
		return C.CString(fmt.Errorf("error closing resource at uri %q: %w", C.GoString(uri), err).Error())
	}

	*str = C.CString(string(b))

	return nil
}

//export WriteString
func WriteString(uri *C.char, alg *C.char, appendFlag C.int, contents *C.char, close C.int) *C.char {

	u := C.GoString(uri)

	writeToResourceOutput, err := grw.WriteToResource(&grw.WriteToResourceInput{
		URI:    u,
		Alg:    C.GoString(alg),
		Append: appendFlag > 0,
	})
	if err != nil {
		return C.CString(fmt.Errorf("error opening resource from uri %q: %w", u, err).Error())
	}
	w := writeToResourceOutput.Writer

	_, err = fmt.Fprint(w, C.GoString(contents))
	if err != nil {
		return C.CString(fmt.Errorf("error writing: %w", err).Error())
	}

	if flusher, ok := w.(interface{ Flush() error }); ok {
		err = flusher.Flush()
		if err != nil {
			return C.CString(fmt.Errorf("error flushing: %w", err).Error())
		}
	}

	if close > 0 {
		err = w.Close()
		if err != nil {
			return C.CString(fmt.Errorf("error closing: %w", err).Error())
		}
	}

	return nil
}

/*

//export ReadBytes
func ReadBytes(uri *C.char, alg *C.char, bytes *unsafe.Pointer, length *C.int) *C.char {

	r, _, err := grw.ReadFromResource(C.GoString(uri), C.GoString(alg), 4096, nil)
	if err != nil {
		return C.CString(fmt.Errorf("error opening resource from uri %q: %w", C.GoString(uri), err).Error())
	}

	b, err := r.ReadAll()
	if err != nil {
		return C.CString(fmt.Errorf("error reading from resource: %w", err).Error())
	}

	*bytes = unsafe.Pointer(&b[0])
	*length = C.int(len(b))

	return nil
}




//export WriteBytes
func WriteBytes(uri *C.char, alg *C.char, appendFlag C.int, bytes unsafe.Pointer, length C.int) *C.char {

	w, err := grw.WriteToResource(C.GoString(uri), C.GoString(alg), appendFlag > 0, nil)
	if err != nil {
		return C.CString(fmt.Errorf("error opening resource from uri %q: %w", C.GoString(uri), err).Error())
	}

	_, err = w.Write(C.GoBytes(bytes, length))
	if err != nil {
		return C.CString(fmt.Errorf("error writing to resource: %w", err).Error())
	}

	return nil
}

*/

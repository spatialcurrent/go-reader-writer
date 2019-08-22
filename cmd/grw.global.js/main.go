// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// grw.global.js is the package for go-reader-writer (GRW) that adds GRW functions to the global scope under the "grw" variable.
//
// In Node, depending on where require is called and the build system used, the functions may need to be required at the top of each module file.
// In a web browser, grw can be made available to the entire web page.
// The functions are defined in the Exports variable in the grwjs package.
//
// Usage
//	// Below is a simple set of examples of how to use this package in a JavaScript application.
//
//	// load functions into global scope
//	// require('./dist/grw.global.min.js);
//
//	// Open a reader for a given uri and compression algorithm.
//	// Returns an object, which can be destructured to the configured reader and error as a string.
//	// If there is no error, then err will be null.
//	var { reader, err } = grw.read(uri, algorithm, options);
//
//	// Open a writer for a given uri and compression algorithm..
//	// Returns an object, which can be destructured to the configured writer and error as a string.
//	// If there is no error, then err will be null.
//	var { writer, err } = grw.write(uri, algorithm, options);
//
// References
//	- https://godoc.org/pkg/github.com/spatialcurrent/go-reader-writer/pkg/grwjs/
//	- https://nodejs.org/api/globals.html#globals_global_objects
//	- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects
package main

import (
	"github.com/gopherjs/gopherjs/js"

	"github.com/spatialcurrent/go-reader-writer/pkg/grwjs"
)

func main() {
	js.Global.Set("grw", grwjs.Exports)
}

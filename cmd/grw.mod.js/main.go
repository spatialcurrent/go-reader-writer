// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// grw.mod.js is the package for go-reader-writer (GRW) that is built as a JavaScript module.
// In modern JavaScript, the module can be imported using destructuring assignment.
// The functions are defined in the Exports variable in the grwjs package.
//
// Usage
//	// Below is a simple set of examples of how to use this package in a JavaScript application.
//
//	// load functions into current scope
//	const { read, write, algorithms, schemes } = require('./dist/grw.global.min.js);
//
//	// Open a reader for a given uri and compression algorithm.
//	// Returns an object, which can be destructured to the configured reader and error as a string.
//	// If there is no error, then err will be null.
//	var { reader, err } = read(uri, algorithm, options);
//
//	// Open a writer for a given uri and compression algorithm..
//	// Returns an object, which can be destructured to the configured writer and error as a string.
//	// If there is no error, then err will be null.
//	var { writer, err } = write(uri, algorithm, options);
//
// References
//	- https://godoc.org/pkg/github.com/spatialcurrent/go-reader-writer/pkg/grwjs/
//	- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment
package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/spatialcurrent/go-reader-writer/pkg/grwjs"
)

func main() {
	jsModuleExports := js.Module.Get("exports")
	for name, value := range grwjs.Exports {
		jsModuleExports.Set(name, value)
	}
}

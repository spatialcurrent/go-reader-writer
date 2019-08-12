// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// grw.js is the Javascript version of go-reader-writer (GRW).
//
package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/pkg/errors"
	"github.com/spatialcurrent/go-reader-writer/pkg/grw"
)

func main() {
	js.Global.Set("grw", map[string]interface{}{
		"version": grw.Version,
		"fetch":   FetchResource,
	})
}

func FetchResource(uri string, alg string, callback func(...interface{}) *js.Object) {

	go func() {

		r, _, err := grw.ReadHTTPFile(uri, alg, false)
		if err != nil {
			callback("", errors.Wrap(err, "error opening resource from uri "+uri).Error())
			return
		}

		b, err := r.ReadAll()
		if err != nil {
			callback("", errors.Wrap(err, "error reading from resource at uri "+uri).Error())
			return
		}

		err = r.Close()
		if err != nil {
			callback("", errors.Wrap(err, "error closing resource at "+uri).Error())
			return
		}

		callback(string(b), nil)
	}()

}

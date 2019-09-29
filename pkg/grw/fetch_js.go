// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// +build js

package grw

import (
	"bytes"
	"io"

	"github.com/gopherjs/gopherjs/js"
	"github.com/pkg/errors"
	"honnef.co/go/js/xhr"
)

// Fetch sends a GET request to the given url and returns a reader for the body, the metadata, and error if any.
func fetch(url string) (io.ReadCloser, *Metadata, error) {
	req := xhr.NewRequest("GET", url)
	req.Timeout = 1000 // one second, in milliseconds
	req.ResponseType = xhr.ArrayBuffer
	//req.OverrideMimeType("text/plain; charset=x-user-defined") // ensure bytes are not mangled
	err := req.Send(nil)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "error opening file from url %q", url)
	}
	b := js.Global.Get("Uint8Array").New(req.Response).Interface().([]byte)
	//fmt.Println(fmt.Sprintf("Data: %x", req.ResponseText))
	return &Reader{Reader: bytes.NewReader(b)}, nil, nil
}

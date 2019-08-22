// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// +build !js

package grw

import (
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func fetch(url string) (io.ReadCloser, *Metadata, error) {
	resp, err := http.Get(url) // #nosec
	if err != nil {
		return nil, nil, errors.Wrapf(err, "error opening file from url %q", url)
	}
	return resp.Body, NewMetadataFromHeader(resp.Header), nil
}

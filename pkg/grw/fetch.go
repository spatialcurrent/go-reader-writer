// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
	"io"
	"net/http"
)

func fetch(url string) (io.ReadCloser, *Metadata, error) {
	resp, err := http.Get(url) // #nosec
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file from url %q: %w", url, err)
	}
	return resp.Body, NewMetadataFromHeader(resp.Header), nil
}

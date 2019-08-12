// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"io"
	"io/ioutil"
	"net/http"
)

import (
	"github.com/pkg/errors"
)

// ReadHTTPFile returns a ByteReadCloser for an object for a web address
// alg may be "bzip2", "gzip", "snappy", "zip", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//  - https://golang.org/pkg/archive/zip/
//
func ReadHTTPFile(uri string, alg string, bufferSize int) (ByteReadCloser, *Metadata, error) {

	resp, err := http.Get(uri) // #nosec
	if err != nil {
		return nil, nil, errors.Wrapf(err, "error opening file from uri %q", uri)
	}

	metadata := NewMetadataFromHeader(resp.Header)

	if alg == AlgorithmZip {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, metadata, errors.Wrapf(err, "error reading bytes from zip-compressed http file at uri %q", uri)
		}
		brc, err := ReadZipBytes(body)
		if err != nil {
			return nil, metadata, errors.Wrapf(err, "error creating reader for zip bytes at uri %q", uri)
		}
		return brc, metadata, nil
	}

	switch alg {
	case AlgorithmBzip2, AlgorithmGzip, AlgorithmSnappy, AlgorithmNone, "":
		r, closers, err := WrapReader(resp.Body, []io.Closer{resp.Body}, alg, bufferSize)
		if err != nil {
			return nil, metadata, errors.Wrapf(err, "error wrapping reader for file at uri %q", uri)
		}
		return &Reader{Reader: r, Closers: closers}, metadata, nil
	}

	return nil, metadata, &ErrUnknownAlgorithm{Algorithm: alg}

}

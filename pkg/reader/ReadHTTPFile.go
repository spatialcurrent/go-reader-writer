// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
	"io/ioutil"
)

// ReadHTTPFile returns a ByteReadCloser for an object for a web address
// alg may be "bzip2", "gzip", "snappy", "zip", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//  - https://golang.org/pkg/archive/zip/
//
func ReadHTTPFile(uri string, alg string, dict []byte, bufferSize int) (*Reader, *Metadata, error) {

	respBody, metadata, err := fetch(uri)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file from uri %q: %w", uri, err)
	}

	/*data, err := ioutil.ReadAll(respBody)
	if err != nil {
		return nil, metadata, fmt.Errorf("error reading bytes from zip-compressed http file at uri %q: %w", uri, err)
	}
	respBody = &Reader{Reader: bytes.NewReader(data), Closer: nil}

	fmt.Println(fmt.Sprintf("Data: % x", string(data)))*/

	if alg == AlgorithmZip {
		body, err := ioutil.ReadAll(respBody)
		if err != nil {
			return nil, metadata, fmt.Errorf("error reading bytes from zip-compressed http file at uri %q: %w", uri, err)
		}
		brc, err := ReadZipBytes(body)
		if err != nil {
			return nil, metadata, fmt.Errorf("error creating reader for zip bytes at uri %q: %w", uri, err)
		}
		return brc, metadata, nil
	}

	switch alg {
	case AlgorithmBzip2, AlgorithmGzip, AlgorithmSnappy, AlgorithmNone, "":
		r, err := WrapReader(respBody, alg, dict, bufferSize)
		if err != nil {
			return nil, metadata, fmt.Errorf("error wrapping reader for file at uri %q: %w", uri, err)
		}
		return &Reader{Reader: r}, metadata, nil
	}

	return nil, metadata, &ErrUnknownAlgorithm{Algorithm: alg}

}

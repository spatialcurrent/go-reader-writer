// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bufio"
	"compress/bzip2"
	"compress/gzip"
	"os"
)

import (
	"github.com/golang/snappy"
	"github.com/pkg/errors"
)

// ReadFromFile returns a ByteReader for a file with a given compression.
// alg may be "snappy", "gzip", or "none."
//
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func ReadFromFile(file *os.File, alg string, cache bool, bufferSize int) (ByteReadCloser, error) {
	switch alg {
	case "snappy":
		reader := &Reader{
			Reader: bufio.NewReaderSize(snappy.NewReader(bufio.NewReaderSize(file, bufferSize)), bufferSize),
			File:   file,
		}
		return reader, nil
	case "gzip":
		gr, err := gzip.NewReader(bufio.NewReaderSize(file, bufferSize))
		if err != nil {
			return nil, errors.Wrap(err, "error creating gzip reader for file \""+file.Name()+"\"")
		}
		reader := &Reader{
			Reader: bufio.NewReaderSize(gr, bufferSize),
			Closer: gr,
			File:   file,
		}
		return reader, nil
	case "bzip2":
		reader := &Reader{
			Reader: bufio.NewReaderSize(bzip2.NewReader(bufio.NewReaderSize(file, bufferSize)), bufferSize),
			File:   file,
		}
		return reader, nil
	case "zip":
		reader, err := ReadZipFile(file.Name(), false)
		if err != nil {
			return nil, errors.Wrap(err, "error creating gzip reader for file \""+file.Name()+"\"")
		}
		return reader, nil
	case "none", "":
		reader := &Reader{
			Reader: bufio.NewReaderSize(file, bufferSize),
			File:   file,
		}
		return reader, nil
	}
	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

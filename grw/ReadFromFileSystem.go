// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

// OpenFile returns a ByteReader for a file with a given compression.
// alg may be "snappy", "gzip", or "none."
//
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func ReadFromFileSystem(uri string, alg string, cache bool, buffer_size int) (ByteReadCloser, error) {
	switch alg {
	case "snappy":
		return ReadSnappyFile(uri, cache, buffer_size)
	case "gzip":
		return ReadGzipFile(uri, cache, buffer_size)
	case "bzip2":
		return ReadBzip2File(uri, cache, buffer_size)
	case "zip":
		return ReadZipFile(uri, cache, buffer_size)
	case "none", "":
		return ReadLocalFile(uri, cache, buffer_size)
	}
	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

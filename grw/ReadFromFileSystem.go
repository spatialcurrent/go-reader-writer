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
func ReadFromFileSystem(path string, alg string, cache bool, buffer_size int) (ByteReadCloser, error) {
	switch alg {
	case "snappy":
		return ReadSnappyFile(path, cache, buffer_size)
	case "gzip":
		return ReadGzipFile(path, cache, buffer_size)
	case "bzip2":
		return ReadBzip2File(path, cache, buffer_size)
	case "zip":
		return ReadZipFile(path, cache)
	case "none", "":
		return ReadLocalFile(path, cache, buffer_size)
	}
	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

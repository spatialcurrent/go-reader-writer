// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

// WriteToFileSystem returns a ByteWriteCloser for a file with a given compression.
// alg may be "snappy", "gzip", or "none."
//
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func WriteToFileSystem(uri string, alg string, flag int, parents bool) (ByteWriteCloser, error) {
	switch alg {
	case "snappy":
		return WriteSnappyFile(uri, flag)
	case "gzip":
		return WriteGzipFile(uri, flag)
	case "bzip2":
		return nil, &ErrWriterNotImplemented{Algorithm: "bzip2"}
	case "zip":
		return nil, &ErrWriterNotImplemented{Algorithm: "zip"}
	case "none", "":
		return WriteLocalFile(uri, flag, parents)
	}
	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

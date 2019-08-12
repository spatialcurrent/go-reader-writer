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
	"io/ioutil"
	"net/url"
	"strings"
	"time"
)

import (
	"github.com/jlaffaye/ftp"
	"github.com/pkg/errors"
)

// ReadFTPFile returns a ByteReadCloser for an object for a ftp, ftps, or sftp address.
// alg may be "bzip2", "gzip", "snappy", "zip", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//  - https://golang.org/pkg/archive/zip/
//
func ReadFTPFile(uri string, alg string, bufferSize int) (ByteReadCloser, error) {

	_, fullpath := SplitUri(uri)

	parts := strings.SplitN(fullpath, "/", 2)
	authority, p := parts[0], parts[1]

	userinfo, host, port := SplitAuthority(authority)
	if len(port) == 0 {
		port = "21"
	}

	conn, err := ftp.Dial(fmt.Sprintf("%s:%s", host, port), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, errors.Wrapf(err, "error opening file from uri %q", uri)
	}

	if len(userinfo) > 0 {
		user := ""
		password := ""
		if strings.Contains(userinfo, ":") {
			parts := strings.SplitN(userinfo, ":", 2)
			if len(parts) == 2 {
				user, err = url.PathUnescape(parts[0])
				if err != nil {
					return nil, errors.Wrapf(err, "error parsing user %q", parts[0])
				}
				password, err = url.PathUnescape(parts[1])
				if err != nil {
					return nil, errors.Wrap(err, "error parsing password")
				}
			}
		} else {
			user = userinfo
		}
		err = conn.Login(user, password)
		if err != nil {
			return nil, errors.Wrapf(err, "error logging in with user %q", user)
		}
	}

	resp, err := conn.Retr(p)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading file from uri %q", uri)
	}

	if alg == AlgorithmZip {
		body, err := ioutil.ReadAll(resp)
		if err != nil {
			return nil, errors.Wrapf(err, "error reading bytes from zip-compressed http fileat uri %q", uri)
		}
		brc, err := ReadZipBytes(body)
		if err != nil {
			return nil, errors.Wrapf(err, "error creating reader for zip bytes at uri %q", uri)
		}
		return brc, nil
	}

	switch alg {
	case AlgorithmBzip2, AlgorithmGzip, AlgorithmSnappy, AlgorithmNone, "":
		r, closers, err := WrapReader(resp, []io.Closer{resp}, alg, bufferSize)
		if err != nil {
			return nil, errors.Wrapf(err, "error wrapping reader for ftp file at uri %q", uri)
		}
		return &Reader{Reader: r, Closers: closers}, nil
	}

	return nil, &ErrUnknownAlgorithm{Algorithm: alg}

}

// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"os"
)

// WriteToResource
func WriteToResource(uri string, alg string, appendFlag bool, s3_client *s3.S3) (ByteWriteCloser, error) {

	if uri == "stdout" {
		return WriteStdout(alg)
	} else if uri == "stderr" {
		return WriteStderr(alg)
	}

	scheme, path := SplitUri(uri)
	switch scheme {
	case "none", "":
		pathExpanded, err := homedir.Expand(path)
		if err != nil {
			return nil, errors.Wrap(err, "Error expanding resource file path "+path)
		}

		flag := 0
		if appendFlag {
			flag = os.O_APPEND | os.O_CREATE | os.O_WRONLY
		} else {
			flag = os.O_CREATE | os.O_WRONLY
		}

		switch alg {
		case "snappy":
			return WriteSnappyFile(pathExpanded, flag)
		case "gzip":
			return WriteGzipFile(pathExpanded, flag)
		case "bzip2":
			return nil, &ErrWriterNotImplemented{Algorithm: "bzip2"}
		case "zip":
			return nil, &ErrWriterNotImplemented{Algorithm: "zip"}
		case "none", "":
			return WriteLocalFile(pathExpanded, flag)
		}
	}

	return nil, &ErrUnknownAlgorithm{Algorithm: alg}
}

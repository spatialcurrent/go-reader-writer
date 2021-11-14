// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package android

import (
	"io"

	"github.com/spatialcurrent/go-reader-writer/pkg/grw"
)

func ReadAll(uri string, alg string) ([]byte, error) {
	readFromResourceOutput, readFromResourceError := grw.ReadFromResource(&grw.ReadFromResourceInput{
		URI:        uri,
		Alg:        alg,
		Dict:       make([]byte, 0),
		BufferSize: -1,
		S3Client:   nil,
		SSHClient:  nil,
		SFTPClient: nil,
		Password:   "",
		PrivateKey: make([]byte, 0),
	})
	if readFromResourceError != nil {
		return nil, readFromResourceError
	}
	data, errReadAll := io.ReadAll(readFromResourceOutput.Reader)
	if errReadAll != nil {
		return nil, errReadAll
	}
	return data, nil
}

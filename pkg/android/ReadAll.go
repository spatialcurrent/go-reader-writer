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
  readFromResourceOutput, err := grw.ReadFromResource(&grw.ReadFromResourceInput{
    URI: uri,
    Alg: alg,
    Dict: make([]byte, 0),
    BufferSize: -1,
    S3Client: nil,
    SSHClient: nil,
    SFTPClient: nil,
    Password: "",
    PrivateKey: make([]byte, 0),
  })
  data, err := io.ReadAll(readFromResourceOutput.Reader)
  if err != nil {
    return nil, err
  }
  return data, nil
}

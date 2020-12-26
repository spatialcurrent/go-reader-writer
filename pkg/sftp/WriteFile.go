// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package sftp

import (
	"fmt"
	"github.com/pkg/sftp"
	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
	"os"
	"strings"
)

func WriteFile(uri string, options ...ClientOption) (*Writer, error) {

	sshClient, err := Dial(uri, options...)
	if err != nil {
		return nil, fmt.Errorf("error creating SSH client: %w", err)
	}

	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return nil, fmt.Errorf("error creating SFTP client: %w", err)
	}

	_, fullpath := splitter.SplitUri(uri)
	parts := strings.SplitN(fullpath, "/", 2)

	file, err := sftpClient.OpenFile(parts[1], os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return NewWriter(file, sftpClient, sshClient), nil

}

// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package ssh2

import (
	"golang.org/x/crypto/ssh"
)

type ClientConfig struct {
	ssh.ClientConfig
}

type ClientOption func(config *ClientConfig) error

type Client struct {
	*ssh.Client
}

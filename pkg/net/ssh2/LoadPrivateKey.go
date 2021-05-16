// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package ssh2

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

type PrivateKey interface {
	ssh.Signer
}

func LoadPrivateKey(path string) (PrivateKey, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading private key from path %q: %w", path, err)
	}
	key, err := ssh.ParsePrivateKey(b)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key from path %q: %w", path, err)
	}
	return key, nil
}

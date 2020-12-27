// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"

	pkgalg "github.com/spatialcurrent/go-reader-writer/pkg/alg"
	"github.com/spatialcurrent/go-reader-writer/pkg/net/sftp2"
	"github.com/spatialcurrent/go-reader-writer/pkg/net/ssh2"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
	"github.com/spatialcurrent/go-reader-writer/pkg/schemes"
	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

type WriteToResourceInput struct {
	Alg        string // compression algorithm
	Append     bool   // append to output resource
	BufferSize int    // buffer size
	Dict       []byte // compression dictionary
	Parents    bool   // automatically create parent directories as necessary
	S3Client   *s3.S3 // AWS S3 Client
	URI        string // uri to write to
	Password   string // password
	PrivateKey []byte // private key
}

type WriteToResourceOutput struct {
	Writer io.WriteCloser
}

// WriteToResource returns a ByteWriteCloser and error, if any.
func WriteToResource(input *WriteToResourceInput) (*WriteToResourceOutput, error) {

	if input.URI == "-" {
		w, err := WrapWriter(os.Stdout, input.Alg, input.Dict, 0)
		if err != nil {
			return nil, fmt.Errorf("error wrapping device %q: %w", input.URI, err)
		}
		return &WriteToResourceOutput{Writer: w}, nil
	}

	scheme, path := splitter.SplitUri(input.URI)
	switch scheme {
	case schemes.SchemeSFTP:
		options := []ssh2.ClientOption{}
		if input.PrivateKey != nil && len(input.PrivateKey) > 0 {
			privateKey, err := ssh.ParsePrivateKey(input.PrivateKey)
			if err != nil {
				return nil, fmt.Errorf("error parsing private key from path %q: %w", path, err)
			}
			options = append(options, func(config *ssh2.ClientConfig) error {
				config.Auth = []ssh.AuthMethod{
					ssh.PublicKeys(privateKey),
				}
				return nil
			})
		} else if len(input.Password) > 0 {
			options = append(options, func(config *ssh2.ClientConfig) error {
				config.Auth = []ssh.AuthMethod{
					ssh.Password(input.Password),
				}
				return nil
			})
		}
		w, err := sftp2.WriteFile(input.URI, options...)
		if err != nil {
			return nil, fmt.Errorf("error creating writer for resource at %q: %w", input.URI, err)
		}
		ww, err := WrapWriter(w, input.Alg, input.Dict, 0)
		if err != nil {
			return nil, fmt.Errorf("error wrapping writer for resource at %q: %w", input.URI, err)
		}
		return &WriteToResourceOutput{Writer: ww}, nil
	case schemes.SchemeFile, "":

		pathExpanded, err := homedir.Expand(path)
		if err != nil {
			return nil, fmt.Errorf("error expanding resource file path %q: %w", path, err)
		}

		flag := 0
		if input.Append {
			flag = os.O_APPEND | os.O_CREATE | os.O_WRONLY
		} else {
			flag = os.O_CREATE | os.O_WRONLY
		}

		if input.Parents {
			err = os.MkdirAll(filepath.Dir(pathExpanded), 0770)
			if err != nil {
				return nil, fmt.Errorf("error creating parent directories: %w", err)
			}
		}

		w, err := WriteToFileSystem(&WriteToFileSystemInput{
			Alg:        input.Alg,
			BufferSize: input.BufferSize,
			Dict:       input.Dict,
			Flag:       flag,
			Parents:    false,
			Path:       pathExpanded,
		})
		if err != nil {
			return nil, fmt.Errorf("error creating writer for resource at %q: %w", input.URI, err)
		}
		return &WriteToResourceOutput{Writer: w}, nil
	}

	return nil, &pkgalg.ErrUnknownAlgorithm{Algorithm: input.Alg}
}

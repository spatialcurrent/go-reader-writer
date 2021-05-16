// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
	"io"
	stdos "os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"

	pkgalg "github.com/spatialcurrent/go-reader-writer/pkg/alg"
	"github.com/spatialcurrent/go-reader-writer/pkg/net/ssh2"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
	"github.com/spatialcurrent/go-reader-writer/pkg/schemes"
	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

type WriteToResourceInput struct {
	ACL        string       // ACL for objects written to AWS s3
	Alg        string       // compression algorithm
	Append     bool         // append to output resource
	BufferSize int          // buffer size
	Dict       []byte       // compression dictionary
	Mode       uint32       // mode of the output file
	Parents    bool         // automatically create parent directories as necessary
	Password   string       // password
	PrivateKey []byte       // private key
	S3Client   *s3.S3       // AWS S3 Client
	SSHClient  *ssh.Client  // SSH Client
	SFTPClient *sftp.Client // SFTP Client
	URI        string       // uri to write to
}

type WriteToResourceOutput struct {
	Writer io.WriteCloser
}

func writeToSFTP(input *WriteToResourceInput) (*WriteToResourceOutput, error) {
	sshClient := input.SSHClient
	if sshClient == nil {
		options := []ssh2.ClientOption{}
		if input.PrivateKey != nil && len(input.PrivateKey) > 0 {
			privateKey, err := ssh.ParsePrivateKey(input.PrivateKey)
			if err != nil {
				return nil, fmt.Errorf("error parsing private SSH key: %w", err)
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
		c, err := ssh2.Dial(input.URI, options...)
		if err != nil {
			return nil, fmt.Errorf("error creating SSH client: %w", err)
		}
		sshClient = c.Client
	}
	sftpClient := input.SFTPClient
	if sftpClient == nil {
		c, err := sftp.NewClient(sshClient)
		if err != nil {
			return nil, fmt.Errorf("error creating SFTP client: %w", err)
		}
		sftpClient = c
	}
	_, fullpath := splitter.SplitURI(input.URI)
	file, err := sftpClient.OpenFile(strings.SplitN(fullpath, "/", 2)[1], os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	// Do not use a SFTP writer, so that the SFTP and SSH connections stay open.
	ww, err := WrapWriter(file, input.Alg, input.Dict, 0)
	if err != nil {
		return nil, fmt.Errorf("error wrapping writer for resource at %q: %w", input.URI, err)
	}
	return &WriteToResourceOutput{
		Writer: &FunctionWriteCloser{
			Writer: func(p []byte) (n int, err error) { return ww.Write(p) },
			Closer: func() error {
				err := ww.Close()
				// After closing the file, attempt to update the desired mode
				if input.Mode != uint32(0) {
					// attempt to change file mode to the desired mode, if error than just continue
					_ = sftpClient.Chmod(file.Name(), stdos.FileMode(input.Mode))
				}
				return err
			},
		},
	}, nil
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

	scheme, path := splitter.SplitURI(input.URI)
	switch scheme {
	case schemes.SchemeSFTP:
		return writeToSFTP(input)
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
			Mode:       input.Mode,
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

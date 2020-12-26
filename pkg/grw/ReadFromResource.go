// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	homedir "github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"

	"github.com/spatialcurrent/go-reader-writer/pkg/io"
	"github.com/spatialcurrent/go-reader-writer/pkg/net/ftp"
	"github.com/spatialcurrent/go-reader-writer/pkg/net/http"
	"github.com/spatialcurrent/go-reader-writer/pkg/net/sftp"
	"github.com/spatialcurrent/go-reader-writer/pkg/net/ssh2"
	"github.com/spatialcurrent/go-reader-writer/pkg/os"
	"github.com/spatialcurrent/go-reader-writer/pkg/schemes"
	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

type ReadFromResourceInput struct {
	URI        string // uri to read from
	Alg        string // compression algorithm
	Dict       []byte // compression dictionary
	BufferSize int    // input reader buffer size
	S3Client   *s3.S3 // AWS S3 Client
	PrivateKey []byte // private key
}

type ReadFromResourceOutput struct {
	Reader   io.ReadCloser
	Metadata *Metadata
}

func fetchRemoteFile(uri string, privateKeyBytes []byte) (io.ReadCloser, error) {
	switch scheme, _ := splitter.SplitUri(uri); scheme {
	case schemes.SchemeFTP:
		return ftp.Fetch(uri)
	case schemes.SchemeSFTP:
		options := []ssh2.ClientOption{}
		if privateKeyBytes != nil && len(privateKeyBytes) > 0 {
			privateKey, err := ssh.ParsePrivateKey(privateKeyBytes)
			if err != nil {
				return nil, fmt.Errorf("error parsing private key: %w", err)
			}
			options = append(options, func(config *ssh2.ClientConfig) error {
				config.Auth = []ssh.AuthMethod{
					ssh.PublicKeys(privateKey),
				}
				return nil
			})
		}
		return sftp.Fetch(uri, options...)
	case schemes.SchemeHTTP, schemes.SchemeHTTPS:
		return http.Fetch(uri)
	}
	return nil, nil
}

func ReadFromResource(input *ReadFromResourceInput) (*ReadFromResourceOutput, error) {

	if input.URI == "-" {
		wr, err := WrapReader(os.Stdin, input.Alg, input.Dict, input.BufferSize)
		if err != nil {
			return nil, fmt.Errorf("error wrapping reader for stdin: %w", err)
		}
		return &ReadFromResourceOutput{Reader: wr, Metadata: nil}, nil
	}

	scheme, path := splitter.SplitUri(input.URI)

	switch scheme {
	case schemes.SchemeFile, "":
		pathExpanded, err := homedir.Expand(path)
		if err != nil {
			return nil, fmt.Errorf("error expanding file path %q: %w", path, err)
		}
		pathCleaned := filepath.Clean(pathExpanded)
		f, err := os.OpenFile(pathCleaned)
		if err != nil {
			return nil, fmt.Errorf("error opening regular file: %w", err)
		}
		wr, err := WrapReader(f, input.Alg, input.Dict, input.BufferSize)
		if err != nil {
			return nil, fmt.Errorf("error wrapping reader for file at uri %q: %w", input.URI, err)
		}
		return &ReadFromResourceOutput{Reader: wr, Metadata: nil}, nil
	case schemes.SchemeFTP, schemes.SchemeSFTP, schemes.SchemeHTTP, schemes.SchemeHTTPS:
		r, err := fetchRemoteFile(input.URI, input.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("error fetching remote file at uri %q: %w", input.URI, err)
		}
		wr, err := WrapReader(r, input.Alg, input.Dict, input.BufferSize)
		if err != nil {
			return nil, fmt.Errorf("error wrapping reader for file at uri %q: %w", input.URI, err)
		}
		return &ReadFromResourceOutput{Reader: wr, Metadata: nil}, nil
	case schemes.SchemeS3:
		i := strings.Index(path, "/")
		if i == -1 {
			return nil, errors.New("path missing bucket")
		}
		r, err := input.S3Client.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(path[0:i]),
			Key:    aws.String(path[i+1:]),
		})
		if err != nil {
			return nil, fmt.Errorf("error fetching file on AWS S3 at uri %q: %w", input.URI, err)
		}
		wr, err := WrapReader(r.Body, input.Alg, input.Dict, input.BufferSize)
		if err != nil {
			return nil, fmt.Errorf("error wrapping reader for file at uri %q: %w", input.URI, err)
		}
		return &ReadFromResourceOutput{Reader: wr, Metadata: NewMetadataFromS3(r)}, nil
	}

	return nil, &schemes.ErrUnknownScheme{Scheme: scheme}
}

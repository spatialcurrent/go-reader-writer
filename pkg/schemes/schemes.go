// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package schemes

const (
	SchemeFile  = "file"
	SchemeFTP   = "ftp"
	SchemeHTTP  = "http"
	SchemeHTTPS = "https"
	SchemeS3    = "s3"
	SchemeSFTP  = "sftp"
)

var (
	Schemes = []string{
		SchemeFile,
		SchemeFTP,
		SchemeHTTP,
		SchemeHTTPS,
		SchemeS3,
		SchemeSFTP,
	}
)

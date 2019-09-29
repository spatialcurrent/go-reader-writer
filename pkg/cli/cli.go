// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package cli

const (
	FlagAWSProfile         string = "aws-profile"
	FlagAWSDefaultRegion   string = "aws-default-region"
	FlagAWSRegion          string = "aws-region"
	FlagAWSAccessKeyID     string = "aws-access-key-id"
	FlagAWSSecretAccessKey string = "aws-secret-access-key"
	FlagAWSSessionToken    string = "aws-session-token"
	FlagInputCompression   string = "input-compression"
	FlagInputDictionary    string = "input-dictionary"
	FlagInputBufferSize    string = "input-buffer-size"
	FlagOutputCompression  string = "output-compression"
	FlagOutputBufferSize   string = "output-buffer-size"
	FlagOutputAppend       string = "output-append"
	FlagOutputOverwrite    string = "output-overwrite"
	FlagOutputDictionary   string = "output-dictionary"
	FlagSplitLines         string = "split-lines"
	FlagVerbose            string = "verbose"

	DefaultBufferSize = 4096

	NumberReplacementCharacter string = "#"
)

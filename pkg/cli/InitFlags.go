// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package cli

import (
	"fmt"

	"github.com/spf13/pflag"
)

func InitFlags(flag *pflag.FlagSet) {
	flag.String(FlagAWSProfile, "", "AWS Profile")
	flag.String(FlagAWSDefaultRegion, "", "AWS Default Region")
	flag.StringP(FlagAWSRegion, "", "", "AWS Region (overrides default region)")
	flag.StringP(FlagAWSAccessKeyID, "", "", "AWS Access Key ID")
	flag.StringP(FlagAWSSecretAccessKey, "", "", "AWS Secret Access Key")
	flag.StringP(FlagAWSSessionToken, "", "", "AWS Session Token")

	flag.StringP(FlagInputCompression, "", "", "the input compression")
	flag.String(FlagInputDictionary, "", "the input dictionary")
	flag.Int(FlagInputBufferSize, DefaultBufferSize, "the input reader buffer size")

	flag.StringP(FlagOutputCompression, "", "", "the output compression")
	flag.String(FlagOutputDictionary, "", "the output dictionary")
	flag.IntP(FlagOutputBufferSize, "b", DefaultBufferSize, "the output writer buffer size")
	flag.BoolP(FlagOutputAppend, "a", false, "append to output files")
	flag.BoolP(FlagOutputOverwrite, "o", false, "overwrite output if it already exists")

	flag.IntP(
		FlagSplitLines,
		"l",
		-1,
		fmt.Sprintf("split output by a number of lines, replaces %q in output uri with file number starting with 1.", NumberReplacementCharacter),
	)

	flag.BoolP(FlagVerbose, "v", false, "verbose output")
}

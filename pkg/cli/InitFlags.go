// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
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

	flag.String(FlagInputCompression, "none", "the input compression")
	flag.String(FlagInputDictionary, "", "the input dictionary")
	flag.Int(FlagInputBufferSize, DefaultBufferSize, "the input reader buffer size")
	flag.String(FlagInputPrivateKey, "", "Use the provided private key to connect to the input.")
	flag.String(FlagInputPassword, "", "Use the provided password to connect to the input.")

	flag.String(FlagOutputCompression, "none", "the output compression")
	flag.String(FlagOutputDictionary, "", "the output dictionary")
	flag.IntP(FlagOutputBufferSize, "b", -1, "The output writer buffer size. The default for stdout is 0.  The default for files is 4096.")

	flag.BoolP(FlagOutputMkdirs, "m", false, "make directories if missing for output file")
	flag.BoolP(FlagOutputAppend, "a", false, "append to output files")
	flag.BoolP(FlagOutputOverwrite, "o", false, "overwrite output if it already exists")
	flag.String(FlagOutputPrivateKey, "", "Use the provided private key to connect to the output.")
	flag.String(FlagOutputPassword, "", "Use the provided password to connect to the output.")

	flag.IntP(
		FlagSplitLines,
		"l",
		-1,
		fmt.Sprintf("split output by a number of lines, replaces %q in output uri with file number starting with 1.", NumberReplacementCharacter),
	)

	flag.Bool(FlagVersion, false, "show version")
	flag.BoolP(FlagVerbose, "v", false, "verbose output")
}

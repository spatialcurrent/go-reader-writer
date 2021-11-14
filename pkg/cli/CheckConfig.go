// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func CheckConfig(args []string, v *viper.Viper) error {

	if len(args) == 0 {
		return fmt.Errorf("missing positional arguments for input and output")
	} else if len(args) == 1 {
		return fmt.Errorf("missing positional argument for output")
	} else if len(args) > 2 {
		return fmt.Errorf("extra positional arguments")
	}

	splitLines := v.GetInt(FlagSplitLines)
	if splitLines > 0 {
		if len(args) < 2 {
			return fmt.Errorf("cannot split by lines when writing to stdout")
		}
		switch args[1] {
		case "stdout", "/dev/stdout", "-":
			return fmt.Errorf("cannot split by lines when writing to stdout")
		}
		outputURI := args[1]
		if !strings.Contains(outputURI, NumberReplacementCharacter) {
			return fmt.Errorf(
				"when splitting by lines, you must include the number replacement character (%q) in the output uri",
				NumberReplacementCharacter,
			)
		}

	}
	return nil
}

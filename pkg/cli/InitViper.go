// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package cli

import (
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func InitViper(flag *pflag.FlagSet) (*viper.Viper, error) {
	v := viper.New()
	err := v.BindPFlags(flag)
	if err != nil {
		return nil, err
	}
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()
	return v, nil
}

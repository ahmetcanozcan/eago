package config

import (
	"fmt"
	"io"

	"github.com/spf13/viper"
)

func readConfig(reader io.Reader, defaultValues map[string]interface{}, fileType string) *viper.Viper {
	v := viper.New()
	v.SetConfigType(fileType)
	for key, val := range defaultValues {
		v.Set(key, val)
	}
	if reader != nil {
		if err := v.ReadConfig(reader); err != nil {
			// TODO: log error with logger system
			fmt.Println("Can not read file", err)
		}
	}
	return v
}

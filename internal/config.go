package internal

import (
	"github.com/spf13/viper"
)

//Configuration is used to store our configuration constants
type Configuration struct{}

//ReadConfig is a utility function which is used to fetch the configuration parameters
func ReadConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {

	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}

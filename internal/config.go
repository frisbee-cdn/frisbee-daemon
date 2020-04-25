package internal

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

//Configuration is used to store our configuration constants
type Configuration struct {
	Server ServerConfiguration

	HashSize int

	Timeout time.Duration
	MaxIdle time.Duration
}

// ServerConfiguration store server configuration details
type ServerConfiguration struct {
	Addr string
	Port uint32
}

func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {

	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(filename)
	v.AddConfigPath("./configs")

	v.AutomaticEnv()
	v.SetConfigType("yml")

	err := v.ReadInConfig()
	return v, err
}

//InitConfiguration is a utility function which is used to fetch
//the configuration parameters
func InitConfiguration(filename string) *Configuration {

	v1, err := readConfig(filename, nil)
	if err != nil {
		panic(fmt.Errorf("Error when reading config: %v\n.", err))
	}

	var serverConf ServerConfiguration

	err = v1.Unmarshal(&serverConf)
	if err != nil {
		panic(fmt.Errorf("Unable to decode configuration into struct %v", err))
	}

	cnf := &Configuration{
		Server: serverConf,
	}

	cnf.HashSize = 160
	cnf.Timeout = 10 * time.Millisecond
	cnf.MaxIdle = 100 * time.Millisecond

	return cnf

}

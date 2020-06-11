package internal

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Mode describes what mode the DHT should operate in
type ModeOpt int

const (

	// MODE_AUTO determines the mode on which the DHT should operate in, switch the DHT
	// between Client and Server based on network conditions
	MODE_AUTO ModeOpt = iota

	// MODE_CLIENT operates the DHT as a client only, it cannot respond to incoming queries
	MODE_CLIENT

	// MODE_SERVER operates the DHT as a server, it ca both send and respond to queries
	MODE_SERVER

	// MODE_AUTO_SERVER operates in the same way as MODE_AUTO, but acts as a server when reachability is unknown
	MODE_AUTO_SERVER
)

//Configuration is used to store our configuration constants
type Configuration struct {
	Server ServerConfiguration

	BucketSize int
	HashSize   int

	mode ModeOpt

	Timeout time.Duration
	MaxIdle time.Duration
}

// ServerConfiguration store server configuration details
type ServerConfiguration struct {
	Addr string
	Port uint32
}

// Defaults represents the default configuration parameters
// TODO: Refactor this into a functins (e.g) GetDefaults() *Configuration {}
var Defaults *Configuration = &Configuration{
	Server: ServerConfiguration{
		Addr: "localhost",
		Port: 8080,
	},
	BucketSize: 20,
	HashSize:   160,
	mode:       MODE_AUTO,
	Timeout:    10 * time.Millisecond,
	MaxIdle:    100 * time.Millisecond,
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

	cnf.BucketSize = 20
	cnf.HashSize = cnf.BucketSize * 8
	cnf.Timeout = 10 * time.Millisecond
	cnf.MaxIdle = 100 * time.Millisecond

	return cnf

}

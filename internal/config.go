package internal

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

//Configuration is used to store our configuration constants
type Configuration struct {
	Server ServerConfiguration

	BucketSize int
	HashSize   int

	IsBootstrap           bool
	DefaultBootstrapPeers []ServerConfiguration

	ParallelismDegree uint32
	Timeout           time.Duration
	MaxIdle           time.Duration

	RepublishTimeout time.Duration
	RefreshTimeout   time.Duration
	ReplicateTimeout time.Duration
	ExpireTimeout    time.Duration
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

	DefaultBootstrapPeers: []ServerConfiguration{
		{Addr: "127.0.0.1", Port: 5001},
	},
	BucketSize: 20,
	HashSize:   160,

	IsBootstrap:       false,
	ParallelismDegree: 3,

	Timeout: 10 * time.Millisecond,
	MaxIdle: 100 * time.Millisecond,

	RepublishTimeout: 86400 * time.Second,
	RefreshTimeout:   3600 * time.Second,
	ReplicateTimeout: 3600 * time.Second,
	ExpireTimeout:    86400 * time.Second,
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

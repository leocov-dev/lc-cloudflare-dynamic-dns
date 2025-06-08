package config

import (
	"github.com/spf13/viper"
)

var (
	C Config
)

type Config struct {
	Cloudflare *Cloudflare `yaml:"cloudflare"`
	Update     *Update     `yaml:"update"`
	err        error
}

func (c *Config) AssertConfigSet() error {
	return c.err
}

type Cloudflare struct {
	ApiToken string `yaml:"apiToken"`
}
type Update struct {
	Name string `yaml:"name"`
	TTL  int    `yaml:"ttl"`
}

func ReadConfigFile(explicitConfig string) {
	if explicitConfig == "" {
		// search for a file named config.yaml
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		// in these directories
		viper.AddConfigPath("/etc/lc-cf-dns")
		viper.AddConfigPath("/etc/cloudflare-dynamic-dns")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("./")
	} else {
		viper.SetConfigFile(explicitConfig)
	}

	err := viper.ReadInConfig()
	if err != nil {
		C.err = err
		return
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		C.err = err
		return
	}
}

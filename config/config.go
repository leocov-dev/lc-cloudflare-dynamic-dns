package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	C Config
)

type Config struct {
	Cloudflare *Cloudflare `yaml:"cloudflare"`
	Update     *Update     `yaml:"update"`
}

type Cloudflare struct {
	ApiToken string `yaml:"apiToken"`
}
type Update struct {
	Name string `yaml:"name"`
	TTL  int    `yaml:"ttl"`
}

func ReadConfigFile() {
	// search for a file named config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// in these directories
	viper.AddConfigPath("/etc/lc-cf-dns")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

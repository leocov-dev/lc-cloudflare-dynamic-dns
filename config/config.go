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
	viper.SetConfigFile("config.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/cf-dns")

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

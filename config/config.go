package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	AppConfig Config
)

type Config struct {
	Cloudflare *Cloudflare
	Update     *Update
}

type Cloudflare struct {
	ApiToken string
}
type Update struct {
	Name     string
	TTL      int
	Interval time.Duration
}

func NewCloudflare(v *viper.Viper) *Cloudflare {
	return &Cloudflare{
		ApiToken: v.GetString("apiToken"),
	}
}

func NewUpdate(v *viper.Viper) *Update {
	return &Update{
		Name:     v.GetString("name"),
		TTL:      v.GetInt("ttl"),
		Interval: v.GetDuration("interval"),
	}
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

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

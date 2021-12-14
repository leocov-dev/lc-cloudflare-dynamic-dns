package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"lc-cloudflare-dynamic-dns/config"
)

var (
	rootCmd = &cobra.Command{
		Use: config.Name,
		Long: fmt.Sprintf("%s %s\nUpdate Cloudflare Dynamic DNS",
			config.Name,
			config.Version),
	}
)

func init() {
	config.ReadConfigFile()

	rootCmd.PersistentFlags().BoolVarP(&config.IsDebug, "debug", "d", false, "Print additional debug and informational logs.")
	_ = rootCmd.PersistentFlags().MarkHidden("debug")

	rootCmd.AddCommand(updateCmd)
}

func Execute() {
	_ = rootCmd.Execute()
}

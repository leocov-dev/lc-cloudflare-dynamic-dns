package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"lc-cloudflare-dynamic-dns/config"
	"os"
	"path/filepath"
)

var (
	exePath, _ = os.Executable()
	exeName    = filepath.Base(exePath)

	rootCmd = &cobra.Command{
		Use: exeName,
		Long: fmt.Sprintf("%s %s\nUpdate Cloudflare Dynamic DNS",
			exeName,
			config.Version),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return config.C.AssertConfigSet()
		},
	}

	explicitConfigFilePath string
)

func init() {
	cobra.OnInitialize(func() {
		config.ReadConfigFile(explicitConfigFilePath)
	})

	rootCmd.PersistentFlags().BoolVarP(&config.IsDebug, "debug", "d", false, "Print additional debug and informational logs.")
	_ = rootCmd.PersistentFlags().MarkHidden("debug")

	rootCmd.PersistentFlags().StringVarP(
		&explicitConfigFilePath,
		"configFile", "c",
		"",
		"The configuration file may be explicitly set otherwise it will be inferred if possible (must be yaml)",
	)

	rootCmd.AddCommand(updateCmd)
}

func Execute() {
	_ = rootCmd.Execute()
}

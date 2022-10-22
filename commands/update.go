package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"lc-cloudflare-dynamic-dns/config"
	"lc-cloudflare-dynamic-dns/internal/cloudflare"
	"lc-cloudflare-dynamic-dns/internal/misc"
	"os"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update Dynamic DNS",
		Run:   runUpdate,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return config.C.AssertConfigSet()
		},
	}
)

func init() {

}

func runUpdate(cmd *cobra.Command, args []string) {
	ip := misc.GetOutboundIP()
	fmt.Printf("Updating DDNS record for: %s => %s\n", config.C.Update.Name, ip)

	api := cloudflare.NewApiClient(config.C.Cloudflare.ApiToken)

	err := api.VerifyAuthToken()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad token: %s\n", err)
		os.Exit(1)
	}

	err = api.DoUpdate(config.C.Update.Name, ip.String(), config.C.Update.TTL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "update error: %s\n", err)
		os.Exit(1)
	}
}

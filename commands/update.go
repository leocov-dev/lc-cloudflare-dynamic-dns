package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"lc-cloudflare-dynamic-dns/config"
	"lc-cloudflare-dynamic-dns/internal/cloudflare"
	"lc-cloudflare-dynamic-dns/internal/misc"
)

var (
	updateCmd = &cobra.Command{
		Use:     "update",
		Short:   "Update Dynamic DNS",
		Aliases: []string{"u"},
		Run:     runUpdate,
	}
)

func init() {

}

func runUpdate(cmd *cobra.Command, args []string) {
	ip := misc.GetOutboundIP()
	fmt.Printf("Updating DDNS record for: %s => %s\n", config.C.Update.Name, ip)

	api := cloudflare.NewApiClient()

	err := api.VerifyAuthToken()
	if err != nil {
		fmt.Println("Bad token: ", err)
	}

	err = api.DoUpdate(config.C.Update.Name, ip.String(), config.C.Update.TTL)
	if err != nil {
		fmt.Println("update error: ", err)
	}
}

package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"lc-cloudflare-dynamic-dns/config"
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
	fmt.Println("doing update!")

	ip := misc.GetOutboundIP()
	fmt.Printf("Update DDNS record for: %s => %s\n", config.AppConfig.Update.Name, ip)
}

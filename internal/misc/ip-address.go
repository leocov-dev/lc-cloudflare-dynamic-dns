package misc

import (
	externalIp "github.com/glendc/go-external-ip"
	"net"
	"os"
)

func GetOutboundIP() net.IP {
	consensus := externalIp.DefaultConsensus(nil, nil)
	ip, err := consensus.ExternalIP()
	if err != nil {
		os.Exit(1)
	}

	return ip
}

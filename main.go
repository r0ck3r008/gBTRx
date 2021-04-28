package main

import (
	"flag"
)

func main() {
	// Input the cmdline arguments
	var common *string = flag.String("common", "Common.cfg", "The Path to Common configuration file")
	var peerInfo *string = flag.String("peerinfo", "PeerInfo.cfg", "The path to PeerInfo configuration file")
	flag.Parse()

	// Parse the peerInfo
	var peers []peerT
	parsePeerInfo(*peerInfo, peers)

	// Parse the Common
	var commonCfg commonT
	parseCommonCfg(*common, &commonCfg)
}

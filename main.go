package main

import (
	"flag"

	"github.com/r0ck3r008/gBTRx/cnfParse"
)

func main() {
	// Input the cmdline arguments
	var common *string = flag.String("common", "Common.cfg", "The Path to Common configuration file")
	var peerInfo *string = flag.String("peerinfo", "PeerInfo.cfg", "The path to PeerInfo configuration file")
	flag.Parse()

	// Parse the peerInfo
	var peers []cnfParse.PeerCfgT
	cnfParse.ParsePeerInfo(*peerInfo, peers)

	// Parse the Common
	var commonCfg cnfParse.CommonCfgT
	cnfParse.ParseCommonCfg(*common, &commonCfg)
}

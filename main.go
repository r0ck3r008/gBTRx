package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type peerT struct {
	id    uint32
	hName string
	port  uint32
	self  bool
}

type commonT struct {
	nPrefNbrs uint32
	uChoke    uint32
	uChokeOp  uint32
	fName     string
	fSz       uint32
	pcSz      uint32
}

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

func fileRead(fName string, res [][]string) {
	file, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		res = append(res, items)
	}
}

func strTou32Safe(str string) uint32 {
	var ret int64
	var err error
	if ret, err = strconv.ParseInt(str, 10, 32); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	return uint32(ret)
}

func parsePeerInfo(peerinfo string, res []peerT) {
	var peers [][]string
	fileRead(peerinfo, peers)

	for _, peerLine := range peers {
		peer := peerT{
			id:    strTou32Safe(peerLine[0]),
			hName: peerLine[1],
			port:  strTou32Safe(peerLine[2]),
			self:  strTou32Safe(peerLine[3]) != 0,
		}

		res = append(res, peer)
	}
}

func parseCommonCfg(common string, ret *commonT) {
	var lines [][]string
	fileRead(common, lines)

	(*ret) = commonT{
		nPrefNbrs: strTou32Safe(lines[0][1]),
		uChoke:    strTou32Safe(lines[1][1]),
		uChokeOp:  strTou32Safe(lines[2][1]),
		fName:     lines[3][1],
		fSz:       strTou32Safe(lines[4][1]),
		pcSz:      strTou32Safe(lines[4][1]),
	}
}

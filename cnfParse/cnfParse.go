package cnfParse

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func ParsePeerInfo(peerinfo string, res []PeerT) {
	var peers [][]string
	fileRead(peerinfo, peers)

	for _, peerLine := range peers {
		peer := PeerT{
			Id:    strTou32Safe(peerLine[0]),
			HName: peerLine[1],
			Port:  strTou32Safe(peerLine[2]),
			Self:  strTou32Safe(peerLine[3]) != 0,
		}

		res = append(res, peer)
	}
}

func ParseCommonCfg(common string, ret *CommonT) {
	var lines [][]string
	fileRead(common, lines)

	(*ret) = CommonT{
		NPrefNbrs: strTou32Safe(lines[0][1]),
		UChoke:    strTou32Safe(lines[1][1]),
		UChokeOp:  strTou32Safe(lines[2][1]),
		FName:     lines[3][1],
		FSz:       strTou32Safe(lines[4][1]),
		PcSz:      strTou32Safe(lines[4][1]),
	}
}
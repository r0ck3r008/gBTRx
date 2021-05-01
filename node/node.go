package node

import (
	"fmt"
	"net"
	"sync"

	"github.com/r0ck3r008/gBTRx/cnfParse"
	"github.com/r0ck3r008/gBTRx/utils"
)

func NewNode(peerId uint32, peerinfo []cnfParse.PeerCfgT) NodeT {
	var wg sync.WaitGroup
	// For each peer that is not itself, create a new tcp connection
	ret := NodeT{
		cliMap: make(map[uint32]*net.TCPConn),
		lSock:  nil,
		wg:     &wg,
	}
	var tcpAddr *net.TCPAddr
	var err error
	var selfCfg cnfParse.PeerCfgT

	for _, peerCfg := range peerinfo {
		if peerCfg.Id == peerId {
			selfCfg = peerCfg
			break
		}
		var conn *net.TCPConn

		if tcpAddr, err = net.ResolveTCPAddr("tcp",
			peerCfg.HName+fmt.Sprintf(":%d", peerCfg.Port)); err != nil {
			utils.ErrExit(err, "Resolve TCP")
		}

		if conn, err = net.DialTCP("tcp", nil, tcpAddr); err != nil {
			utils.ErrExit(err, "Dial TCP")
		}

		peer := PeerT{
			peerId: peerCfg.Id,
			conn:   conn,
			server: false,
		}
		wg.Add(1)
		// NOTE: It is safe to copy contents of PeerT to the new
		// GoRoutine since these are thread safe
		go func(wg *sync.WaitGroup, p PeerT) {
			CliLoop()
			wg.Done()
		}(&wg, peer)
	}

	if tcpAddr, err = net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", selfCfg.Port)); err != nil {
		utils.ErrExit(err, "Resolve TCP")
	}

	var lSock *net.TCPListener
	if lSock, err = net.ListenTCP("tcp", tcpAddr); err != nil {
		utils.ErrExit(err, "Listen TCP")
	}
	ret.lSock = lSock

	return ret
}

func (node *NodeT) AcceptLoop() {
	for {
		conn, err := (*node.lSock).AcceptTCP()
		if err != nil {
			utils.ErrExit(err, "Accept TCP")
		}

		node.wg.Add(1)
		go func(wg *sync.WaitGroup, conn *net.TCPConn) {
			CliLoop()
			wg.Done()
		}(node.wg, conn)
	}
}

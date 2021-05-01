package node

import (
	"net"
	"sync"
)

type NodeT struct {
	cliMap map[uint32]*net.TCPConn
	lSock  *net.TCPListener
	wg     *sync.WaitGroup
}

type PeerT struct {
	peerId uint32
	conn   *net.TCPConn
	server bool
}

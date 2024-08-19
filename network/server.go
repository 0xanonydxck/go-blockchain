package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transport []Transport
}

type Server struct {
	ServerOpts
	rpcCh  chan RPC
	quitCh chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransport()
	ticker := time.NewTicker(5 * time.Second)

FREE:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("from: %v , payload: %v\n", rpc.From, string(rpc.Payload))
		case <-s.quitCh:
			break FREE
		case <-ticker.C:
			fmt.Println("do stuff every x seconds")
		}
	}

	fmt.Println("Server stopped")
}

func (s *Server) initTransport() {
	for _, tr := range s.Transport {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}

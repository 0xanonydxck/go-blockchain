package main

import (
	"time"

	"github.com/dxckboi/go-blockchain/network"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")
	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello, Local!"))
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			trLocal.SendMessage(trRemote.Addr(), []byte("Hello, Remote!"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transport: []network.Transport{
			trLocal,
			trRemote,
		},
	}

	serv := network.NewServer(opts)
	serv.Start()
}

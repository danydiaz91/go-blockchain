package main

import (
	"github.com/danydiaz91/go-blockchain/network"
	"time"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("hello world"))
			time.Sleep(time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal, trRemote},
	}

	s := network.NewServer(opts)
	s.Start()
}
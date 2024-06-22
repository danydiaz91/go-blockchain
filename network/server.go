package network

import (
	"fmt"
	"time"
)

type options struct {
	Transports []Transport
}

type Option func(options) options

func WithTransports(transports []Transport) Option {
	return func(serverOpts options) options {
		serverOpts.Transports = transports
		return serverOpts
	}
}

type Server struct {
	options
	rpcCh  chan RPC
	quitCh chan struct{}
}

func NewServer(opts ...Option) *Server {
	serverOpts := options{}
	for _, opt := range opts {
		serverOpts = opt(serverOpts)
	}

	return &Server{
		options: serverOpts,
		rpcCh:   make(chan RPC),
		quitCh:  make(chan struct{}),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(5 * time.Second)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			fmt.Println("do stuff every 5 seconds")
		}
	}

	fmt.Println("server shutdown")
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}

package main

import (
	"context"
	"log"
	"net"
	rpc_ping "github.com/YTakahashii/hello-grpc-go/interfaces/rpc/ping"
	"google.golang.org/grpc"
)

type server struct{}

const port = ":3001"

// implements
type IPingImplement interface {
	Ping(ctx context.Context, p *rpc_ping.Empty) (*rpc_ping.Pong, error)
}

type pingImplement struct {
}

// NewPingImplement is ...
func NewPingImplement() IPingImplement {
	return &pingImplement{}
}

func (imp *pingImplement) Ping(ctx context.Context, p *rpc_ping.Empty) (*rpc_ping.Pong, error) {
	res := rpc_ping.Pong{}
	res.Reply = "Pong"
	return &res, nil
}

func main() {
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	var grpcServer = grpc.NewServer()
	grpcServer.Serve(listenPort)
}

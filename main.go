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

// SayHello implements helloworld.GreeterServer
func (s *server) Ping(ctx context.Context, in *rpc_ping.Empty) (*rpc_ping.Pong, error) {
	return &rpc_ping.Pong{Reply: "Hello"}, nil
}

func main() {
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	rpc_ping.RegisterCheckServer(grpcServer, &server{})
	grpcServer.Serve(listenPort)
}

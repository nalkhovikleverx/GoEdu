package main

import (
	"GoEdu/proto/ping"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (s *server) Ping(ctx context.Context, p *ping.PingRequest) (*ping.PingResponse, error) {
	println("REQUEST MESSAGE:", p.Message)
	return &ping.PingResponse{Message: "Pong"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	ping.RegisterPingServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

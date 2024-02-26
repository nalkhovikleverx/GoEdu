package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"GoEdu/proto/ping"
)

type server struct {
	ping.UnimplementedPingServiceServer
}

func (s *server) Ping(ctx context.Context, p *ping.PingRequest) (*ping.PingResponse, error) {
	log.Println("REQUEST MESSAGE:", p.Message)
	return &ping.PingResponse{Message: "Pong"}, nil
}

func main() {
	lis, errLis := net.Listen("tcp", "127.0.0.1:8000")
	if errLis != nil {
		log.Fatal(errLis)
	}
	grpcServer := grpc.NewServer()
	ping.RegisterPingServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

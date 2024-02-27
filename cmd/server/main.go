package main

import (
	"GoEdu/proto/ping"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	ping.UnimplementedPingServiceServer
}

func (s *server) Ping(ctx context.Context, p *ping.PingRequest) (*ping.PingResponse, error) {
	log.Println("REQUEST MESSAGE:", p.Message)
	return &ping.PingResponse{Message: "Pong"}, nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

}

func run() error {
	lis, errLis := net.Listen("tcp", "0.0.0.0:8000")
	if errLis != nil {
		return errLis
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()
	ping.RegisterPingServiceServer(grpcServer, &server{})

	return grpcServer.Serve(lis)
}

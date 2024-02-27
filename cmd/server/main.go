package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

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
	errChan := make(chan error)
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	lis, errLis := net.Listen("tcp", "0.0.0.0:8000")
	if errLis != nil {
		log.Fatal(errLis)
	}
	grpcServer := grpc.NewServer()
	ping.RegisterPingServiceServer(grpcServer, &server{})

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			errChan <- err
		}
	}()
	defer func() {
		grpcServer.GracefulStop()
	}()

	select {
	case err := <-errChan:
		log.Println("Fatal error:", err)
	case <-stopChan:
	}

}

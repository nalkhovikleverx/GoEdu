package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"GoEdu/proto/ping"
)

func main() {
	conn, errConn := grpc.Dial("0.0.0.0:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if errConn != nil {
		log.Fatal(errConn)
	}

	client := ping.NewPingServiceClient(conn)
	response, errResp := client.Ping(context.Background(), &ping.PingRequest{Message: "Ping"})
	if errResp != nil {
		log.Fatal(errResp)
	}
	defer conn.Close()
	log.Println("Response: ", response.Message)
}

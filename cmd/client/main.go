package main

import (
	"GoEdu/proto/ping"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := ping.NewPingServiceClient(conn)
	response, err := client.Ping(context.Background(), &ping.PingRequest{Message: "Ping"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response: ", response.Message)
}

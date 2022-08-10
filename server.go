package main

import (
	"log"
	"net"

	"github.com/kkothule/grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("failed to listen on port 9000: %v", err)
	}
	log.Print("started server")
	grpcServer := grpc.NewServer()
	s := chat.Server{}
	chat.RegisterPingServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve on port 9000: %v", err)
	}
}

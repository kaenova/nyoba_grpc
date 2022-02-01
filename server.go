package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kaenova/nyoba_grpc/service/chat"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Registering all services
	c := chat.Server{}
	chat.RegisterChatServiceServer(grpcServer, &c)

	// go func() { // Kalau masih ada service lain yang mau dinyalakan
	// 	if err := grpcServer.Serve(lis); err != nil {
	// 		log.Fatalf("failed to serve: %s", err)
	// 	}
	// }()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

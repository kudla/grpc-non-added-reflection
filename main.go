package main

import (
	"fmt"
	"log"
	"net"

	protos "grpc-non-added-reflection/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port = 3333

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Have a reference to the definition
	log.Printf("Listen on %v port with imported service desc %v\n", port, protos.Test_ServiceDesc)

	var opts = []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)

	// protos.RegisterTestServer(grpcServer, &protos.UnimplementedTestServer{})

	reflection.Register(grpcServer)

	grpcServer.Serve(lis)
}

package main

import (
	"grocerylist/internals"
	ls "grocerylist/protos/listserver"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	PORT = ":8000"
)

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to Listen : %v", err)
	}
	s := grpc.NewServer()
	ls.RegisterListServiceServer(s, &internals.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve :%v", err)
	}

}

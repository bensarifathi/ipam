package main

import (
	"log"
	"net"
	"os"

	"github.com/BENSARI-Fathi/ipam/grpc"
)

func main() {

	socketFile := "/tmp/my-ipam.sock"

	if err := os.RemoveAll(socketFile); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", socketFile)
	if err != nil {
		log.Fatalf("Error while binding socket %s", err.Error())
	}
	defer l.Close()
	s := grpc.NewGrpcServer()
	defer s.Stop()
	log.Printf("Grpc server listenning on Unix %s", socketFile)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve %v\n", err)
	}
}

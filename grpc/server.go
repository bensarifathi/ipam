package grpc

import (
	"github.com/BENSARI-Fathi/ipam/v1/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterIpamServer(s, &Server{})
	return s
}

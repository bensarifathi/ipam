package grpc

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/BENSARI-Fathi/ipam/utils"
	"github.com/BENSARI-Fathi/ipam/v1/pb"
)

type Server struct {
	pb.UnimplementedIpamServer
}

func (s *Server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	lock := &sync.Mutex{}
	log.Println("Add Rpc was called !")
	// extract the fields
	subnet := req.GetSubnet()
	gw := req.GetGateway()
	cID := req.GetContainerId()
	// verify if the creds are correct
	_, ipNet, err := net.ParseCIDR(subnet)
	if err != nil {
		log.Fatalf("Invalid creds %s\n", err.Error())
	}
	if isValid := net.ParseIP(gw); isValid == nil {
		log.Fatalf("Invalid creds %s\n", gw)
	}

	lock.Lock()
	defer lock.Unlock()

	// load the mapped data
	ipToid, err := utils.LoadMap()
	if err != nil {
		return nil, err
	}
	// alocate the first free ip

	ipFromFile, err := utils.LoadIpfile()
	if err != nil {
		return nil, err
	}

	newIp := utils.AlocateIP(ipFromFile, net.ParseIP(gw))
	mask, _ := ipNet.Mask.Size()
	// map the allocated ip to the container id and save it to a file
	ipToid[cID] = newIp
	if err = utils.DumpMap(&ipToid); err != nil {
		return nil, err
	}
	// prepare the response
	resp := &pb.AddResponse{
		PodIp:   newIp,
		Gateway: gw,
		NetMask: int32(mask),
	}

	return resp, nil
}

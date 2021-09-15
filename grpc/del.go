package grpc

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/BENSARI-Fathi/ipam/utils"
	"github.com/BENSARI-Fathi/ipam/v1/pb"
)

func (s *Server) Del(ctx context.Context, req *pb.DelRequest) (*pb.DelResponse, error) {
	lock := &sync.Mutex{}
	cID := req.GetContainerId()
	// extract the ip addr from the file
	myMap, err := utils.LoadMap()
	if err != nil {
		return nil, err
	}
	ip, ok := myMap[cID]
	if !ok {
		ioutil.WriteFile("/var/log/error.log", []byte(fmt.Sprintf("err key (%s) not found", cID)), 0644)
		return &pb.DelResponse{}, nil
	}
	lock.Lock()
	foutput, err := utils.LoadIpfile()
	if err != nil {
		return nil, err
	}
	// release the ip addr
	newFreeIP := utils.FreeIp{IP: &ip}
	foutput.FreeIps = append(foutput.FreeIps, newFreeIP)
	// update the files
	if err = utils.DumpIpfile(foutput); err != nil {
		return nil, err
	}
	delete(myMap, cID)
	if err = utils.DumpMap(&myMap); err != nil {
		return nil, err
	}
	lock.Unlock()
	// send the response
	return &pb.DelResponse{}, nil
}

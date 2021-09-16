package utils

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/c-robinson/iplib"
)

var IpConfFile = "/tmp/ip.json"
var contIDtoIP = "/tmp/ip-id-map.json"

type IpParse struct {
	LastUsedIp *string  `json:"last_used_ip"`
	FreeIps    []FreeIp `json:"free_ips,omitempty"`
}

type FreeIp struct {
	IP *string `json:"ip"`
}

func AlocateIP(ipFile *IpParse, gw net.IP) string {
	if ipFile.FreeIps != nil {
		ip := ipFile.FreeIps[0].IP
		newFreeIP := &IpParse{
			LastUsedIp: ipFile.LastUsedIp,
			FreeIps:    ipFile.FreeIps[1:],
		}
		out, err := json.Marshal(newFreeIP)
		if err != nil {
			log.Fatalf("Error while marshaling FreeIpList %s\n", err.Error())
		}
		if err := ioutil.WriteFile(IpConfFile, out, 0644); err != nil {
			log.Fatalf("Error while updating config file %s\n", err.Error())
		}
		return *ip
	}

	if ipFile.LastUsedIp != nil {
		lastIP := net.ParseIP(*ipFile.LastUsedIp)
		newIP := iplib.IncrementIP4By(lastIP, 1)
		newIpStr := fmt.Sprintf("%v", newIP)
		// update the conf file
		ipFile.LastUsedIp = &newIpStr
		// marshall the data
		out, err := json.Marshal(ipFile)
		if err != nil {
			log.Fatalf("Error while updating ipam config file %s", err.Error())
		}
		// wirte to file
		ioutil.WriteFile(IpConfFile, out, 0644)
		return newIpStr
	}
	newIP := iplib.IncrementIP4By(gw, 1)
	newIpStr := fmt.Sprintf("%v", newIP)
	// update the conf file
	ipFile.LastUsedIp = &newIpStr
	// marshall the data
	out, err := json.Marshal(ipFile)
	if err != nil {
		log.Fatalf("Error while updating ipam config file %s", err.Error())
	}
	// wirte to file
	ioutil.WriteFile(IpConfFile, out, 0644)
	return newIpStr

}

func RemoveIP(ipFile *IpParse, podIp string) error {
	freeIP := ipFile.FreeIps
	newFreeIP := FreeIp{IP: &podIp}
	freeIP = append(freeIP, newFreeIP)
	newConf := ipFile
	newConf.FreeIps = freeIP
	out, err := json.Marshal(newConf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(IpConfFile, out, fs.FileMode(os.O_WRONLY))
}

func LoadMap() (map[string]string, error) {
	myMap := make(map[string]string)
	f, err := os.Open(contIDtoIP)
	if err != nil {
		if os.IsNotExist(err) {
			// will be hit only if the file is not created yet.
			return myMap, nil
		}
		return nil, err
	}
	defer f.Close()
	output, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(output, &myMap); err != nil {
		return nil, err
	}
	return myMap, nil
}

func DumpMap(newMap *map[string]string) error {
	in, err := json.Marshal(newMap)
	if err != nil {
		return err
	}
	if err = os.WriteFile(contIDtoIP, in, 0644); err != nil {
		return err
	}
	return nil
}

func LoadIpfile() (*IpParse, error) {
	ips := &IpParse{}
	f, err := os.OpenFile(IpConfFile, os.O_RDWR, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			// will be hit only if the file is not created yet.
			return ips, nil
		}
		return nil, err
	}
	defer f.Close()
	output, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(output, ips); err != nil {
		return nil, err
	}
	return ips, nil
}

func DumpIpfile(ips *IpParse) error {
	in, err := json.Marshal(ips)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(IpConfFile, in, 0644); err != nil {
		return err
	}
	return nil
}

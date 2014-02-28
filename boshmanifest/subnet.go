package boshmanifest

import "net"
import "fmt"

type subnet struct {
	Range           string                 `json:"range"`
	Reserved        []string               `json:"reserved"`
	Static          []string               `json:"static"`
	Gateway         string                 `json:"gateway"`
	Dns             []string               `json:"dns"`
	CloudProperties networkCloudProperties `json:"cloud_properties"`
}

type networkCloudProperties struct {
	Name string `json:"name"`
}

func newSubnet(
	cidr string,
	dns []string,
	networkName string) (s *subnet, err error) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return
	}

	reserved := fmt.Sprintf("%s - %s", usableIp(ipNet.IP, 2), usableIp(ipNet.IP, 10))
	static := fmt.Sprintf("%s - %s", usableIp(ipNet.IP, 11), usableIp(ipNet.IP, 20))

	s = &subnet{
		ipNet.String(),
		[]string{reserved},
		[]string{static},
		usableIp(ipNet.IP, 1).String(),
		dns,
		networkCloudProperties{networkName},
	}

	return
}

func usableIp(networkIp net.IP, index int) (ip net.IP) {
	ip = net.IP{
		networkIp[0] + byte((index&(255<<24))>>24),
		networkIp[1] + byte((index&(255<<16))>>16),
		networkIp[2] + byte((index&(255<<8))>>8),
		networkIp[3] + byte(index&255),
	}

	return
}

package service

import (
	"fmt"
	"net"
	"os"
)

func IpAddressGen(subnet string) []string {
	ipAddress, ipNet, err := net.ParseCIDR(subnet)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var ipAddresses []string
	for ipAddress := ipAddress.Mask(ipNet.Mask); ipNet.Contains(ipAddress); inc(ipAddress) {
		ipAddresses = append(ipAddresses, ipAddress.String())
	}

	return ipAddresses

}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

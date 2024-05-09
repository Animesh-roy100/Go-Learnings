package main

import (
	"fmt"
	"net"
)

func main() {
	// Define the CIDR subnet
	cidr := "2405:0200:0808:2272::/64"

	// Parse the CIDR
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Iterate over all addresses in the subnet
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		fmt.Println(ip)
	}
}

// Function to increment IPv6 address
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

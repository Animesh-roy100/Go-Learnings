package utils

import (
	"net"
	"strings"
)

// IsValidCIDR checks if the given string is a valid CIDR notation
func IsValidCIDR(cidr string) bool {
	// Split the CIDR into IP and prefix
	parts := strings.Split(cidr, "/")
	if len(parts) != 2 {
		return false
	}

	// Parse the IP address
	ip := net.ParseIP(parts[0])
	if ip == nil {
		return false
	}

	// Check if it's IPv4 or IPv6
	if ip.To4() != nil {
		// IPv4
		return isValidIPv4CIDR(cidr)
	}

	// IPv6
	return isValidIPv6CIDR(cidr)
}

// isValidIPv4CIDR checks if the given string is a valid IPv4 CIDR notation
func isValidIPv4CIDR(cidr string) bool {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}

	// Ensure it's actually an IPv4 address
	if ipNet.IP.To4() == nil {
		return false
	}

	// Check if the prefix is between 0 and 32
	_, bits := ipNet.Mask.Size()
	return bits == 32
}

// isValidIPv6CIDR checks if the given string is a valid IPv6 CIDR notation
func isValidIPv6CIDR(cidr string) bool {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}

	// Ensure it's actually an IPv6 address
	if ipNet.IP.To4() != nil {
		return false
	}

	// Check if the prefix is between 0 and 128
	_, bits := ipNet.Mask.Size()
	return bits == 128
}

// Helper function to filter IPv4 DNS servers
func FilterIPv4DNSServers(servers []string) []string {
	var ipv4Servers []string
	for _, server := range servers {
		ip := net.ParseIP(server)
		if ip != nil && ip.To4() != nil {
			ipv4Servers = append(ipv4Servers, server)
		}
	}
	return ipv4Servers
}

// Helper function to filter IPv6 DNS servers
func FilterIPv6DNSServers(servers []string) []string {
	var ipv6Servers []string
	for _, server := range servers {
		ip := net.ParseIP(server)
		if ip != nil && ip.To4() == nil {
			ipv6Servers = append(ipv6Servers, server)
		}
	}
	return ipv6Servers
}

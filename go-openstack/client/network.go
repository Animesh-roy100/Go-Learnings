package client

import (
	models "openstack/model"
	"openstack/utils"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
)

func ListNetworks() ([]networks.Network, error) {
	allPages, err := networks.List(NetworkClient, networks.ListOpts{}).AllPages()
	if err != nil {
		return []networks.Network{}, err
	}

	allNetworks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		return []networks.Network{}, err
	}

	return allNetworks, nil
}

func CreateNetwork(req models.CreateNetworkRequest) (*networks.Network, error) {
	createOpts := networks.CreateOpts{
		Name:         req.NetworkName,
		AdminStateUp: gophercloud.Enabled,
	}

	network, err := networks.Create(NetworkClient, createOpts).Extract()
	if err != nil {
		return nil, err
	}

	for _, subnet := range req.Subnets {
		ipVersion := gophercloud.IPv6
		var dnsServers []string
		if subnet.IsIPv4 {
			ipVersion = gophercloud.IPv4

			dnsServers = utils.FilterIPv4DNSServers(req.DNSNames)
		} else {
			dnsServers = utils.FilterIPv6DNSServers(req.DNSNames)
		}

		subnetOpts := subnets.CreateOpts{
			NetworkID:      network.ID,
			CIDR:           subnet.CIDR,
			IPVersion:      ipVersion,
			Name:           subnet.SubnetName,
			EnableDHCP:     gophercloud.Enabled,
			DNSNameservers: dnsServers,
		}

		_, err = subnets.Create(NetworkClient, subnetOpts).Extract()
		if err != nil {
			// If subnet creation fails, delete the network and return error
			networks.Delete(NetworkClient, network.ID)
			return nil, err
		}
	}

	return network, nil
}

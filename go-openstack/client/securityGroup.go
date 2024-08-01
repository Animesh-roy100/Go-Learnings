package client

import (
	models "openstack/model"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
)

func CreateSecurityGroup(req models.CreateSecurityGroupRequest) (*groups.SecGroup, error) {
	createOpts := groups.CreateOpts{
		Name: req.SecurityGroupName,
	}

	secGroup, err := groups.Create(NetworkClient, createOpts).Extract()
	if err != nil {
		return nil, err
	}

	return secGroup, nil
}

func AddSSHRule(groupID string) (*rules.SecGroupRule, error) {
	createOpts := rules.CreateOpts{
		Direction:      "ingress",
		EtherType:      "IPv4",
		SecGroupID:     groupID,
		PortRangeMin:   22,
		PortRangeMax:   22,
		Protocol:       "tcp",
		RemoteIPPrefix: "0.0.0.0/0",
	}

	return rules.Create(NetworkClient, createOpts).Extract()
}

func AddEgressICMPRule(groupID string) (*rules.SecGroupRule, error) {
	createOpts := rules.CreateOpts{
		Direction:      "egress",
		EtherType:      "IPv4",
		SecGroupID:     groupID,
		Protocol:       "icmp",
		RemoteIPPrefix: "0.0.0.0/0",
	}

	return rules.Create(NetworkClient, createOpts).Extract()
}

package models

type CreateNetworkRequest struct {
	NetworkName string   `json:"networkName"`
	Subnets     []Subnet `json:"subnets"`
	DNSNames    []string `json:"dnsNames"`
}

type Subnet struct {
	SubnetName string `json:"subnetName"`
	CIDR       string `json:"cidr"`
	IsIPv4     bool   `json:"isIPv4"`
}

type CreateRouterRequest struct {
	RouterName        string `json:"routerName"`
	AdminStateUp      bool   `json:"adminStateUp"`
	ExternalNetworkID string `json:"externalNetworkId"`
}

type AddRouterInterfaceRequest struct {
	SubnetID string `json:"subnetId" binding:"required"`
}

type CreateSecurityGroupRequest struct {
	SecurityGroupName string `json:"securityGroupName"`
}

package client

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

var NetworkClient *gophercloud.ServiceClient

func InitNetworkClient() {
	// Set up authentication
	// ProviderClient: holds the authentication token and can be used to build a ServerClient
	authOption := gophercloud.AuthScope{
		ProjectName: "admin",
		DomainName:  "Default",
	}

	providerClient, err := openstack.AuthenticatedClient(gophercloud.AuthOptions{
		IdentityEndpoint: "http://192.168.64.20/identity/v3",
		Username:         "admin",
		Password:         "secret",
		DomainName:       "Default",
		AllowReauth:      true,
		Scope:            &authOption,
	})
	if err != nil {
		panic(err)
	}

	NetworkClient = &gophercloud.ServiceClient{
		ProviderClient: providerClient,
		Endpoint:       "http://192.168.64.20:9696/networking/v2.0/",
	}
}

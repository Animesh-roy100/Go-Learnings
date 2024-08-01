package client

import (
	models "openstack/model"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
)

func CreateRouter(req models.CreateRouterRequest) (*routers.Router, error) {
	createOpts := routers.CreateOpts{
		Name:         req.RouterName,
		AdminStateUp: &req.AdminStateUp,
	}

	if req.ExternalNetworkID != "" {
		createOpts.GatewayInfo = &routers.GatewayInfo{
			NetworkID: req.ExternalNetworkID,
		}
	}

	router, err := routers.Create(NetworkClient, createOpts).Extract()
	if err != nil {
		return nil, err
	}

	return router, nil
}

func AddRouterInterface(routerId string, req models.AddRouterInterfaceRequest) (*routers.InterfaceInfo, error) {
	addOpts := routers.AddInterfaceOpts{
		SubnetID: req.SubnetID,
	}

	interfaceInfo, err := routers.AddInterface(NetworkClient, routerId, addOpts).Extract()
	if err != nil {
		return nil, err
	}

	return interfaceInfo, nil
}

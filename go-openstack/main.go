package main

import (
	"openstack/client"
	"openstack/initializers"
	"openstack/router"
)

func init() {
	initializers.LoadEnv()
	client.InitNetworkClient()
}

func main() {
	// Run the router
	router.Run()
}

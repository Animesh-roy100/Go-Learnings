package handler

import (
	"net/http"
	"openstack/client"
	models "openstack/model"
	"openstack/utils"

	"github.com/gin-gonic/gin"
)

func ListNetworksHandler(c *gin.Context) {
	allNetworks, err := client.ListNetworks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, allNetworks)
}

func CreateNetworkHandler(c *gin.Context) {
	var req models.CreateNetworkRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate CIDRs
	for _, subnet := range req.Subnets {
		if !utils.IsValidCIDR(subnet.CIDR) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CIDR: " + subnet.CIDR})
			return
		}
	}

	network, err := client.CreateNetwork(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, network)
}

package handler

import (
	"net/http"
	"openstack/client"
	models "openstack/model"

	"github.com/gin-gonic/gin"
)

func CreateSecurityGroupHandler(c *gin.Context) {
	var req models.CreateSecurityGroupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	securityGroup, err := client.CreateSecurityGroup(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create security group " + err.Error()})
		return
	}

	// Add SSH rule
	_, err = client.AddSSHRule(securityGroup.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add SSH rule: " + err.Error()})
		return
	}

	// Add Egress ICMP rule
	_, err = client.AddEgressICMPRule(securityGroup.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add Egress ICMP rule: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, securityGroup)
}

package handler

import (
	"net/http"
	"openstack/client"
	models "openstack/model"

	"github.com/gin-gonic/gin"
)

func CreateRouterHandler(c *gin.Context) {
	var req models.CreateRouterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	router, err := client.CreateRouter(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, router)
}

func AddRouterInterfaceHandler(c *gin.Context) {
	routerId := c.Param("id")
	var req models.AddRouterInterfaceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	interfaceInfo, err := client.AddRouterInterface(routerId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, interfaceInfo)
}

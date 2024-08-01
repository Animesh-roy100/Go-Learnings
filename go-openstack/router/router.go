package router

import (
	"fmt"
	handler "openstack/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.New()

	networkGroup := router.Group("/network")
	networkGroup.GET("", handler.ListNetworksHandler)
	networkGroup.POST("", handler.CreateNetworkHandler)
	// networkGroup.GET("/:id", handler.GetNetworkHandler)
	// networkGroup.PUT("/:id", handler.UpdateNetworkHandler)
	// networkGroup.DELETE("/:id", handler.DeleteNetworkHandler)

	routerGroup := router.Group("/router")
	// routerGroup.GET("", handler.ListRoutersHandler)
	routerGroup.POST("", handler.CreateRouterHandler)
	routerGroup.POST("/:id/add-interface", handler.AddRouterInterfaceHandler)

	securityGroup := router.Group("/security-groups")
	securityGroup.POST("", handler.CreateSecurityGroupHandler)

	port := os.Getenv("PORT")
	fmt.Println("Server listening on :" + port)
	if err := router.Run(":" + port); err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}

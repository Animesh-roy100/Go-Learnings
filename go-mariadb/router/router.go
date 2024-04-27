package router

import (
	"fmt"
	"os"

	"github.com/Animesh-roy100/go-mariadb/controllers"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.New()

	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	port := os.Getenv("PORT")
	fmt.Println("Server listening on :" + port)
	if err := router.Run(":" + port); err != nil {
		fmt.Printf("Error on starting server: %v", err)
	}
}

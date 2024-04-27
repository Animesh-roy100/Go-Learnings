package db

import (
	"fmt"
	"os"

	"github.com/Animesh-roy100/go-mariadb/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBUser *gorm.DB

func LoadMariaDB() {
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")

	dsn := username + ":" + password + "@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DBUser, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to MariaDB:", err)
		return
	}

	if DBUser != nil {
		DBUser.AutoMigrate(&models.User{})
		fmt.Println("Successfully connected to MariaDB and migrated schema")
	} else {
		fmt.Println("Failed to connect to MariaDB")
	}
}

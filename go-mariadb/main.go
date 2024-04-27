package main

import (
	"github.com/Animesh-roy100/go-mariadb/db"
	"github.com/Animesh-roy100/go-mariadb/initializers"
	"github.com/Animesh-roy100/go-mariadb/router"
)

func main() {
	initializers.LoadEnv()
	db.LoadMariaDB()
	router.Run()
}

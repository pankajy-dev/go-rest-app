package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pankajy-dev/go-rest-app/db"
	"github.com/pankajy-dev/go-rest-app/routes"
)

func main() {
	db.InitDb()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}

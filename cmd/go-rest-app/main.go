package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pankajy-dev/go-rest-app/internal/config"
	"github.com/pankajy-dev/go-rest-app/internal/db"
	"github.com/pankajy-dev/go-rest-app/internal/handlers"
)

func main() {
	db.InitDb()
	server := gin.Default()

	handlers.RegisterRoutes(server)

	server.Run(config.ServerPort())
}

package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pankajy-dev/go-rest-app/internal/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // events/1

	// Either add middleware individually or as a group
	//server.POST("/events", middleware.Authenticate, createEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)

	authenticated.POST("/events", middleware.Authenticate, createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelEventRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pankajy-dev/go-rest-app/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // events/1

	// Either add middleware individually or as a group
	//server.POST("/events", middlewares.Authenticate, createEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", middlewares.Authenticate, createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", deleteEventRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}

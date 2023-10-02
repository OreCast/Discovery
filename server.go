package main

import (
	"fmt"
	"log"

	authz "github.com/OreCast/Authz/auth"
	"github.com/gin-gonic/gin"
)

// examples: https://go.dev/doc/tutorial/web-service-gin

// helper function to setup our server router
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// GET routes
	r.GET("/sites", SitesHandler)

	// all POST methods ahould be authorized
	authorized := r.Group("/")
	authorized.Use(authz.TokenMiddleware(Config.AuthzClientId, Config.Verbose))
	{
		authorized.POST("/sites", SitesPostHandler)
	}

	return r
}

func Server(configFile string) {
	r := setupRouter()
	sport := fmt.Sprintf(":%d", Config.Port)
	log.Printf("Start HTTP server %s", sport)
	r.Run(sport)
}

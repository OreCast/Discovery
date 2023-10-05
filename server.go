package main

import (
	"fmt"
	"log"

	authz "github.com/OreCast/common/authz"
	oreConfig "github.com/OreCast/common/config"
	oreMongo "github.com/OreCast/common/mongo"
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
	authorized.Use(authz.TokenMiddleware(oreConfig.Config.Authz.ClientId, oreConfig.Config.Discovery.Verbose))
	{
		authorized.POST("/sites", SitesPostHandler)
		authorized.DELETE("/site/:site", SiteDeleteHandler)
	}

	return r
}

func Server() {
	// init MongoDB
	oreMongo.InitMongoDB(oreConfig.Config.MetaData.DBUri)

	// setup web router and start the service
	r := setupRouter()
	sport := fmt.Sprintf(":%d", oreConfig.Config.Discovery.Port)
	log.Printf("Start HTTP server %s", sport)
	r.Run(sport)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// examples: https://go.dev/doc/tutorial/web-service-gin

// helper function to setup our server router
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/sites", func(c *gin.Context) {
		data := _sites
		c.AsciiJSON(http.StatusOK, data)
	})
	r.POST("/sites", func(c *gin.Context) {
		var site Site
		err := c.BindJSON(&site)
		if err == nil {
			found := false
			for _, s := range _sites {
				if s.Name == site.Name && s.URL == site.URL {
					found = true
					break
				}
			}
			if !found {
				_sites = append(_sites, site)
			}
			c.JSON(200, gin.H{"status": "ok"})
		} else {
			c.JSON(400, gin.H{"status": "fail", "error": err.Error()})
		}
	})
	return r
}

func Server(configFile string) {
	r := setupRouter()
	sport := fmt.Sprintf(":%d", Config.Port)
	log.Printf("Start HTTP server %s", sport)
	r.Run(sport)
}

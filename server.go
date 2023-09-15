package main

import (
	"fmt"
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
			_sites = append(_sites, site)
			c.JSON(200, gin.H{"status": "ok"})
		} else {
			c.JSON(400, gin.H{"status": "fail", "error": err.Error()})
		}
	})
	return r
}

func Server(configFile string) {
	r := setupRouter()
	r.Run(fmt.Sprintf(":%d", Config.Port))
}

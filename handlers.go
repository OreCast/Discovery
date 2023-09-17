package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SitesHandler provides access to GET /sites end-point
func SitesHandler(c *gin.Context) {
	data := _sites
	c.AsciiJSON(http.StatusOK, data)
}

// SitesPostHandler provides access to POST /sites end-point
func SitesPostHandler(c *gin.Context) {
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
}

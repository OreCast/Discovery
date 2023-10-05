package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SiteParams represents URI storage params in /meta/:site end-point
type SiteParams struct {
	Site string `uri:"site" binding:"required"`
}

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
		// upsert into MongoDB
		site.mongoUpsert("Name")
		c.JSON(200, gin.H{"status": "ok"})
	} else {
		c.JSON(400, gin.H{"status": "fail", "error": err.Error()})
	}
}

// SiteDeleteHandler provides access to Delete /site/:mid end-point
func SiteDeleteHandler(c *gin.Context) {
	var params SiteParams
	if err := c.ShouldBindUri(&params); err == nil {
		var sites []Site
		for _, site := range _sites {
			if site.Name != params.Site {
				sites = append(sites, site)
				// remove record from MongoDB
				site.mongoRemove()
			}
		}
		if len(_sites) == len(sites) {
			// record was not found
			msg := fmt.Sprintf("site %s was not found in Discovery service", params.Site)
			c.JSON(400, gin.H{"status": "fail", "error": msg})
			return
		}
		_sites = sites
		c.JSON(200, gin.H{"status": "ok"})
	} else {
		c.JSON(400, gin.H{"status": "fail", "error": err.Error()})
	}
}

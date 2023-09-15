package main

// Site object
type Site struct {
	Name string `json:"name" binding:"required"`
	URL  string `json:"url" binding:"required"`
}

// global site list we keep in server
var _sites []Site

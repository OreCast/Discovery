package main

import (
	oreConfig "github.com/OreCast/common/config"
	oreMongo "github.com/OreCast/common/mongo"
	"gopkg.in/mgo.v2/bson"
)

// Site object
type Site struct {
	Name         string `json:"name" binding:"required"`
	URL          string `json:"url" binding:"required"`
	Endpoint     string `json:"endpoint" binding:"required"`
	AccessKey    string `json:"access_key" binding:"required"`
	AccessSecret string `json:"access_secret" binding:"required"`
	UseSSL       bool   `json:"use_ssl"`
	Description  string `json:"description"`
}

// Record converts Site to MongoDB record
func (s *Site) Record() oreMongo.Record {
	rec := make(oreMongo.Record)
	rec["name"] = s.Name
	rec["url"] = s.URL
	rec["endpoint"] = s.Endpoint
	rec["access_key"] = s.AccessKey
	rec["access_secret"] = s.AccessSecret
	rec["use_ssl"] = s.UseSSL
	rec["description"] = s.Description
	return rec
}

// insert Site record to MongoDB
func (s *Site) mongoInsert() {
	var records []oreMongo.Record
	records = append(records, s.Record())
	oreMongo.Insert(
		oreConfig.Config.Discovery.MongoDB.DBName,
		oreConfig.Config.Discovery.MongoDB.DBColl,
		records)
}

// upsert Site record to MongoDB by using given key
func (s *Site) mongoUpsert(key string) {
	var records []oreMongo.Record
	records = append(records, s.Record())
	oreMongo.Upsert(
		oreConfig.Config.Discovery.MongoDB.DBName,
		oreConfig.Config.Discovery.MongoDB.DBColl,
		key,
		records)
}

// remove Site record from MongoDB
func (s *Site) mongoRemove() {
	spec := bson.M{"name": s.Name}
	oreMongo.Remove(
		oreConfig.Config.Discovery.MongoDB.DBName,
		oreConfig.Config.Discovery.MongoDB.DBColl,
		spec)
}

// global site list we keep in server
var _sites []Site

package main

import (
	"github.com/cyrus28214/comments-area/config"
	"github.com/cyrus28214/comments-area/database"
	"github.com/cyrus28214/comments-area/logging"
	"github.com/cyrus28214/comments-area/server"
	"github.com/gin-gonic/gin"
)

func main() {
	configFile := "./config.yaml"
	cfg := config.GetConfig(configFile)

	log := logging.GetLogger(cfg.Logging)
	log.Infof("Load config from %s", configFile)

	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Panicf("Failed to connect to database: %s", err)
	}
	log.Infof("Connected to database")

	err = database.Migrate(db)
	if err != nil {
		log.Panicf("Failed to migrate database: %s", err)
	}
	log.Info("Database migrated")

	router := gin.New()

	s := server.NewServer(cfg.Server, log, db, router)
	s.Start()
}

package app

import (
	"task/internal/config"
	"task/pkg/database/mongodb"
	"task/pkg/logger"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)
		return
	}

	mongoClient, err := mongodb.NewClient(cfg.Mongo.URI, cfg.Mongo.User, cfg.Mongo.Password)
	if err != nil {
		logger.Error(err)
		return
	}

	_ = mongoClient.Database(cfg.Mongo.Name)

}

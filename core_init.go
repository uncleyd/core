package core

import (
	"core/config"
	"core/db"
	"core/logger"
)

func Init(cfgPath string) {
	config.LoadConfig(cfgPath)
	logger.NewZap(config.Get().Logger[0])
	db.Init()

	logger.Sugar.Info("config & logger & db is init complete!")
}

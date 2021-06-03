package core

import (
	"github.com/uncleyd/core/config"
	"github.com/uncleyd/core/db"
	"github.com/uncleyd/core/logger"
)

func Init(cfgPath string) {
	config.LoadConfig(cfgPath)
	logger.NewZap(config.Get().Logger[0])
	db.Init()

	logger.Sugar.Info("config & logger & db is init complete!")
}

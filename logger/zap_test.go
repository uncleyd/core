package logger

import (
	"testing"
	"time"
	"core/config"
)

func TestZapLog(t *testing.T) {
	NewZap(&config.LoggerConfig{
		Path:         "E:\\git_sh\\za_proxy\\admin-panel\\logs\\",
		Suffix:       ".log",
		Level:        "debug",
		IsWriteFile:  true,
		MaxAge:       7,
		RotationHour: 2,
	})

	go func() {
		for {
			Sugar.Debug("debug=====================")
			Sugar.Info("info---------------------------------------")
			Sugar.Warn("warn---------------------------------------")
			time.Sleep(1000)
		}
	}()
}

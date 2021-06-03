package logger

import (
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/uncleyd/core/config"
	"github.com/uncleyd/core/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"path"
	"time"
)

var Sugar *zap.SugaredLogger

func NewZap(cfg *config.LoggerConfig) {
	if cfg.IsWriteFile {
		NewFileZap(
			cfg.Path,
			cfg.GetFileName(),
			cfg.MaxAge,
			cfg.RotationHour,
			cfg.GetLevel(),
		)
	} else {
		newConsoleZap(cfg)
	}
}

func newConsoleZap(cfg *config.LoggerConfig) {
	dev := zap.NewDevelopmentConfig()
	dev.Level = zap.NewAtomicLevelAt(cfg.GetLevel())
	dev.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	builder, err := dev.Build()
	if err != nil {
		panic(err)
	}

	Sugar = builder.Sugar()
}

// NewFileZap create a file output rule for log
func NewFileZap(filePath, fileName string, maxAge, rotationHour int, level zapcore.Level) {
	utils.CheckPath(filePath)

	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:     "ts",
		LevelKey:    "level",
		MessageKey:  "msg",
		CallerKey:   "file",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	Sugar = zap.New(zapcore.NewTee(
		zapcore.NewCore(encoder,
			zapcore.AddSync(getWriter(filePath, fileName, maxAge, rotationHour)),
			zap.NewAtomicLevelAt(level),
		),
	), zap.AddCaller()).Sugar()
}

func getWriter(filePath, fileName string, maxAge, rotationHour int) io.Writer {
	f := path.Join(filePath, fileName)

	hook, err := rotateLogs.New(
		f+".%Y%m%d%H%M",
		rotateLogs.WithLinkName(f),
		rotateLogs.WithMaxAge(time.Hour*24*time.Duration(maxAge)),
		rotateLogs.WithRotationTime(time.Minute*time.Duration(rotationHour)), //Hour
	)

	if err != nil {
		panic(err)
	}
	return hook
}

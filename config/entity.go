package config

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type AdminConfig struct {
	Debug     bool            `json:"debug"`
	WebServer string          `json:"web_server"`
	Limit     *Limit          `json:"limit"`
	Logger    []*LoggerConfig `json:"logger"`
	Gin       *GinConfig      `json:"gin"`
	MySql     []*MySqlConfig  `json:"mysql"`
	Redis     []*RedisConfig  `json:"redis"`
	Rpcx      RpcxConfig      `json:"rpcx"`
}

func (a *AdminConfig) IsDebug() bool {
	return a.Debug
}

type Limit struct {
	Num   string `json:"num"`
	Clean int    `json:"clean"`
}

type MySqlConfig struct {
	Enable         bool   `json:"enable"`
	GroupId        string `json:"groupId"`
	Id             string `json:"id"`
	DbName         string `json:"dbName"`
	Host           string `json:"host"`
	UserName       string `json:"userName"`
	Password       string `json:"password"`
	MaxIdleConnect int    `json:"maxIdleConnect"`
	MaXOpenConnect int    `json:"maxOpenConnect"`
	LogMode        bool   `json:"logMode"`
}

type RedisConfig struct {
	Id             string `json:"id"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	Password       string `json:"password"`
	Index          int    `json:"index"`
	MaxConnect     int    `json:"maxConnect"`
	MaxIdleConnect int    `json:"maxIdleConnect"`
	MinIdleConnect int    `json:"minIdleConnect"`
}

type GinConfig struct {
	View               string `json:"view"`
	StaticRelativePath string `json:"staticRelativePath"`
	StaticRootPath     string `json:"staticRootPath"`
	Favicon            string `json:"favicon"`
	FaviconPath        string `json:"faviconPath"`
	URL                string `json:"url"`
	Port               int    `json:"port"`
	IsApi              bool   `json:"is_api"`
}

type LoggerConfig struct {
	Path         string `json:"path"`         // 文件保存路径
	Suffix       string `json:"suffix"`       // 文件保存格式
	Level        string `json:"level"`        // 日志等级
	IsWriteFile  bool   `json:"isWriteFile"`  // 是否写文件
	MaxAge       int    `json:"maxAge"`       // 日志保存天数
	RotationHour int    `json:"rotationHour"` // 日志更新时间
}

func (l *LoggerConfig) GetFileName() string {
	return l.Level + l.Suffix
}

func (l *LoggerConfig) GetLevel() zapcore.Level {
	switch strings.ToLower(l.Level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

type RpcxConfig struct {
	ConsulAddr string `json:"consulAddr"` // consul地址
	BasePath   string `json:"basePath"`   // 路径
}

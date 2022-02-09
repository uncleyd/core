package server

import (
	"context"
	"fmt"
	//"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/uncleyd/core/config"
	"github.com/uncleyd/core/logger"
	"github.com/uncleyd/core/utils"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var sugar *zap.SugaredLogger

type WebServer struct {
	*gin.Engine
	controllers []IController
}

func New() *WebServer {
	var cfg = config.Get()

	sugar = logger.Sugar

	ws := &WebServer{}

	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
		ws.Engine = gin.Default()
		//pprof.Register(ws.Engine,"dev/pprof")
		//ws.Engine.Use(RedisLimit())
	} else {
		gin.SetMode(gin.ReleaseMode)
		ws.Engine = gin.New()
		ws.Engine.Use(gin.Recovery()) //logger.WriteFile(),
		//ws.Engine.Use(RedisLimit())
	}
	if cfg.Limit.Status{
		ws.Engine.Use(RedisLimit())
	}
	

	if !cfg.Gin.IsApi {
		var files []string
		files = utils.GetAllFile(cfg.Gin.View, files)
		if len(files) < 1 {
			panic(fmt.Sprintf("not found template files. path = %s", cfg.Gin.View))
		}

		//for _, v := range files {
		//	sugar.Debugf("load template file = %s", v)
		//}
		ws.Engine.LoadHTMLFiles(files...)

		// resources/view/**/*
		//ws.Engine.LoadHTMLGlob(cfg.Gin.View)
		// "/static", "resources/static/"
		ws.Engine.Static(cfg.Gin.StaticRelativePath, cfg.Gin.StaticRootPath)
		//"favicon.ico", "resources/static/favicon.ico"
		ws.Engine.StaticFile(cfg.Gin.Favicon, cfg.Gin.FaviconPath)
	}

	return ws
}

func (w *WebServer) Run() {
	//init controllers
	for _, v := range w.controllers {
		v.Init(w.Engine)
	}

	sugar.Infof("debug http://localhost:%d", config.Get().Gin.Port)

	http := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Get().Gin.Port),
		Handler:      w.Engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := http.ListenAndServe(); err != nil {
			sugar.Fatal(err)
		}
	}()

	fmt.Printf("http server Running on %s:%d \n", config.Get().Gin.URL, config.Get().Gin.Port)

	exit(http)
}

func exit(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	sugar.Infof("got a signal %s", sig)
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		sugar.Error(err)
	}

	sugar.Infof("------shutdown--------%s", time.Since(now))
}

func (w *WebServer) AddController(controllers ...IController) {
	for _, c := range controllers {
		w.controllers = append(w.controllers, c)
	}
}

func (w *WebServer) Do(f func(s *gin.Engine)) {
	f(w.Engine)
}

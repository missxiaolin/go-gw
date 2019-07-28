package main

import (
	"fmt"
	"github.com/zctod/go-tool/common/util_server"
	"go-gw/bootstrap"
	"go-gw/config"
	"go-gw/web/middleware"
	"go-gw/web/routes"
	"net/http"
	"time"
)

func newApp() *bootstrap.Bootstrapper {
	// 初始化应用
	app := bootstrap.New("官网", "xiaolin")
	app.Bootstrap()
	app.Configure(routes.Configure)

	return app
}

func main()  {
	app := newApp()
	app.Use(middleware.Cors())
	startServer(app)
}

func startServer (b *bootstrap.Bootstrapper)  {
	server := &http.Server{
		Addr:           ":" + config.Cfg.Produce.Port,
		Handler:        b,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	// 平滑退出，先结束所有在执行的任务
	util_server.GracefulExitWeb(server)
}
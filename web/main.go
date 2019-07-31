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

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func newApp() *bootstrap.Bootstrapper {
	// 初始化应用
	app := bootstrap.New("北瑟官网", "xiaolin")
	app.Bootstrap()
	app.Use(middleware.Cors())
	app.Configure(routes.Configure, routes.AdminConfigure, routes.ApiConfigure)

	return app
}

func main()  {
	app := newApp()

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

	//if run, ok := commands[runtime.GOOS]; ok {
	//	cmd := exec.Command(run, "http://localhost:" + config.Cfg.Produce.Port)
	//	_ = cmd.Start()
	//}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	// 平滑退出，先结束所有在执行的任务
	util_server.GracefulExitWeb(server)
}
package main

import (
	"binghambai.com/lowCodePlatform-mobile/app/conn"
	"binghambai.com/lowCodePlatform-mobile/app/middleware"
	"binghambai.com/lowCodePlatform-mobile/app/routers"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type GinEngine func(*gin.Engine)

func main() {
	//强制日志颜色化
	gin.ForceConsoleColor()
	//禁用日志颜色
	//gin.DisableConsoleColor()

	//连接redis
	conn.InitConn()

	//注册所有路由
	router := gin.Default()

	//添加自定义中间件
	router.Use(middleware.CatchError())

	//注册所有路由
	routers.RegisterAllRouters(router)

	//初始化服务器，加载服务器相关配置
	initServer(router)

	var g GinEngine
	g(router)
}

func initServer(router *gin.Engine) {

	srv := &http.Server{
		Addr:           ":8001",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		//服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(timeout); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

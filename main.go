package main

import (
	"fmt"
	"net/http"

	"go-gin-blog/models"
	"go-gin-blog/pkg/logging"
	"go-gin-blog/pkg/redis"
	"go-gin-blog/pkg/setting"
	"go-gin-blog/router"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	redis.Setup()

	router := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	fmt.Println("sssss")
}

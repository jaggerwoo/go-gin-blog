package main

import (
	"fmt"
	"net/http"
	"go-gin-blog/pkg/setting"
	"go-gin-blog/router"
)

func main() {
	router := router.InitRouter()

	s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeout,
        WriteTimeout:   setting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }
	s.ListenAndServe()
	fmt.Sprintf(":%d", setting.HTTPPort)
	fmt.Println("sssss")
}
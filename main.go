package main

import (
	"fmt"
	"github.com/chenxiang0010/gin-example/routers"
	"net/http"

	"github.com/chenxiang0010/gin-example/pkg/setting"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

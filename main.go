package main

import (
	"log"
	"net/http"

	"gamenews.niracler.com/monitor/controller"
	"gamenews.niracler.com/monitor/service"
	"gamenews.niracler.com/monitor/setting"
	"github.com/robfig/cron/v3"
)

func main() {
	service.ConnectDB()

	c := cron.New()
	_, _ = c.AddFunc("*/20 * * * *", service.UpdateUserOperation)
	_, _ = c.AddFunc("*/15 * * * *", service.UpdatePVUV)
	c.Start()
	defer c.Stop()

	router := controller.MapRoutes()

	server := &http.Server{
		Addr:           "0.0.0.0:" + setting.HTTPPort,
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}

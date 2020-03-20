package main

import (
	"context"
	"cumtBook/pkg/logging"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"cumtBook/pkg/setting"
	"cumtBook/routers"
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
	log.Println("Listen to: ", setting.HTTPPort)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
			logging.Info("Listen : ", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	logging.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown", err)
		logging.Fatal("Server Shutdown", err)
	}

	log.Println("Server EXiting")
	logging.Info("Server Exiting")
}

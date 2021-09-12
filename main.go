package main

import (
	"WSServer/config"
	"WSServer/handler"
	"WSServer/logger"
	"WSServer/server"
	"context"
	"fmt"
	"github.com/subchen/go-log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.Init()
	logger.Init()
	server.Init()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	httpServer := &http.Server{
		Handler:      handler.NewHandler(),
		Addr:         fmt.Sprintf(":%d", config.Get().Port),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	defer shutdown(httpServer, ctx)
	go serve(httpServer)

	gs := make(chan os.Signal, 1)
	signal.Notify(gs, syscall.SIGTERM, syscall.SIGINT)
	<-gs

	log.Println("exit")
	os.Exit(0)
}

func serve(s *http.Server) {
	log.Infof("start server: %s", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("cant start server: %v", err)
	}
}

func shutdown(s *http.Server, ctx context.Context) {
	log.Infof("stop server: %s", s.Addr)
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("cant stop server: %v", err)
	}
}

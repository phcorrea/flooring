package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger().Print("FloorFinder is up and running")

	srv := &http.Server{
		Handler:      router(),
		Addr:         fmt.Sprintf(":%s", os.Getenv("HTTP_SERVER_PORT")),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	idleConnectionsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			logger().Printf("HTTP server Shutdown: %v", err)
		}

		close(idleConnectionsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger().Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnectionsClosed

	logger().Print("FloorFinder is down")
}

func logger() *log.Logger {
	prefix := ""
	loggerRef = log.New(os.Stdout, prefix, log.LstdFlags)
	return loggerRef
}

var loggerRef *log.Logger

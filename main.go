package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	logger().Print("FloorFinder is up and running")
	logger().Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("HTTP_SERVER_PORT")), router()))
}

func logger() *log.Logger {
	prefix := ""
	loggerRef = log.New(os.Stdout, prefix, log.LstdFlags)
	return loggerRef
}

var loggerRef *log.Logger

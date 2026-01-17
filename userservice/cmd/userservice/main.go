package main

import (
	"os"
	"os/signal"
	"syscall"
	"userservice/internal/app"
	"userservice/internal/config"
	"userservice/pkg/logger"

	_ "github.com/lib/pq"
)

func main() {
	// LOAD CONFIG
	config := config.MustLoad()

	// SETUP LOGGER
	log := logger.SetupLogger(config.LogConf.Level)

	app := app.NewApp(&config, log)

	go app.Run()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)

	<-quitCh

	app.Stop()
}

package main

import (
	"os"
	"os/signal"
	"syscall"
	"userservice/internal/app"

	_ "github.com/lib/pq"
)

func main() {
	app := app.NewApp()

	go app.Run()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)

	<-quitCh

	app.Stop()
}

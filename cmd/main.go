package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/PepsiKingIV/KeyValueDB/internal/server/app"
	"go.uber.org/zap"
)

func main() {

	logger := zap.Must(zap.NewProduction())
	port := 4012
	application := app.New(logger, port)

	go application.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	application.Stop()
}

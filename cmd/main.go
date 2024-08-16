package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/PepsiKingIV/KeyValueDB/config"
	"github.com/PepsiKingIV/KeyValueDB/internal/server/app"
	"github.com/PepsiKingIV/KeyValueDB/pkg/db"
	"go.uber.org/zap"
)

func main() {
	ctxb := context.Background()
	conf := config.WithDefault(ctxb)
	logger := zap.Must(zap.NewProduction())
	database := db.NewDB()
	application := app.New(database, logger, conf.Port)

	go application.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	application.Stop()
}

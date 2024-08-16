package app

import (
	"fmt"
	"net"

	storeRPC "github.com/PepsiKingIV/KeyValueDB/internal/server/gRPC/store"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	logger     *zap.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(logger *zap.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	storeRPC.Register(gRPCServer)
	return &App{
		logger:     logger,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) Run() error {
	const op = "app.Run"
	a.logger.With(
		zap.String("func name", op),
		zap.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.logger.Info("gRPC server is running", zap.String("server addr", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	const op = "app.Stop"

	a.logger.With(
		zap.String("func name", op),
		zap.Int("port", a.port),
	).Info("stopping gRPC server", zap.Int("part", a.port))

	a.gRPCServer.GracefulStop()
}

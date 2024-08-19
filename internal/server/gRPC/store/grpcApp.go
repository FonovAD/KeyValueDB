package store

import (
	"context"
	"errors"
	"time"

	pb "github.com/PepsiKingIV/KeyValueDB/internal/server/gRPC/interfaces"
	"github.com/PepsiKingIV/KeyValueDB/pkg/db"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	pb.UnimplementedStoreServer
	store  db.Store
	logger *zap.Logger
}

func Register(gRPCServer *grpc.Server, store db.Store, logger *zap.Logger) {
	pb.RegisterStoreServer(gRPCServer, &serverAPI{store: store, logger: logger})
}

func (s *serverAPI) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	start := time.Now()
	if req.GetKey() == "" || req.GetValue() == "" {
		return nil, status.Error(codes.InvalidArgument, "key and value cannot be of zero length")
	}
	err := s.store.Put(req.Key, req.Value)
	if err != nil {
		return nil, status.Error(codes.Unknown, errors.Join(ErrStorePut, err).Error())
	}
	s.logger.Info("Set request", zap.String("Method", "SET"), zap.Duration("ResponseTime(ns)", time.Since(start)))
	return &pb.SetResponse{Status: "1", Key: req.GetKey()}, nil
}

func (s *serverAPI) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	start := time.Now()
	if req.GetKey() == "" {
		return nil, status.Error(codes.InvalidArgument, "key cannot be of zero length")
	}
	value, err := s.store.Get(req.Key)
	switch {
	case err == db.ErrRecordNotFound:
		return nil, status.Error(codes.NotFound, errors.Join(ErrStoreGet, err).Error())
	case err != nil:
		return nil, status.Error(codes.Unknown, errors.Join(ErrStoreGet, err).Error())
	}
	s.logger.Info("Get request", zap.String("Method", "GET"), zap.Duration("ResponseTime(ns)", time.Since(start)))
	return &pb.GetResponse{Key: req.GetKey(), Value: value}, nil
}
func (s *serverAPI) Delete(ctx context.Context, req *pb.DelRequest) (*pb.DelResponse, error) {
	start := time.Now()
	if req.GetKey() == "" {
		return nil, status.Error(codes.InvalidArgument, "key cannot be of zero length")
	}
	err := s.store.Delete(req.GetKey())
	switch {
	case err == db.ErrRecordNotFound:
		return nil, status.Error(codes.NotFound, errors.Join(ErrStoreDel, err).Error())
	case err != nil:
		return nil, status.Error(codes.Unknown, errors.Join(ErrStoreDel, err).Error())
	}
	s.logger.Info("Set request", zap.String("Method", "SET"), zap.Duration("ResponseTime(ns)", time.Since(start)))
	return &pb.DelResponse{Key: req.GetKey(), Status: "1"}, nil
}

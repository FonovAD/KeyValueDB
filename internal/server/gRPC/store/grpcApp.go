package store

import (
	"context"
	"errors"

	pb "github.com/PepsiKingIV/KeyValueDB/internal/server/gRPC/interfaces"
	"github.com/PepsiKingIV/KeyValueDB/pkg/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	pb.UnimplementedStoreServer
	store db.Store
}

func Register(gRPCServer *grpc.Server, store db.Store) {
	pb.RegisterStoreServer(gRPCServer, &serverAPI{store: store})
}

func (s *serverAPI) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	if req.GetKey() == "" || req.GetValue() == "" {
		return nil, status.Error(codes.InvalidArgument, "key and value cannot be of zero length")
	}
	err := s.store.Put(req.Key, req.Value)
	if err != nil {
		return nil, status.Error(codes.Unknown, errors.Join(ErrStorePut, err).Error())
	}
	return &pb.SetResponse{Status: "1", Key: req.GetKey()}, nil
}

func (s *serverAPI) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	if req.GetKey() == "" {
		return nil, status.Error(codes.InvalidArgument, "key cannot be of zero length")
	}
	value, err := s.store.Get(req.Key)
	if err != nil {
		return nil, status.Error(codes.Unknown, errors.Join(ErrStoreGet, err).Error())
	}
	return &pb.GetResponse{Key: req.GetKey(), Value: value}, nil
}
func (s *serverAPI) Delete(ctx context.Context, req *pb.DelRequest) (*pb.DelResponse, error) {
	if req.GetKey() == "" {
		return nil, status.Error(codes.InvalidArgument, "key cannot be of zero length")
	}
	err := s.store.Delete(req.GetKey())
	if err != nil {
		return nil, status.Error(codes.Unknown, errors.Join(ErrStoreDel, err).Error())
	}
	return &pb.DelResponse{Key: req.GetKey(), Status: "1"}, nil
}

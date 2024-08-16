package store

import (
	"context"

	pb "github.com/PepsiKingIV/KeyValueDB/internal/server/gRPC/interfaces"
	"google.golang.org/grpc"
)

type serverAPI struct {
	pb.UnimplementedStoreServer
}

func Register(gRPCServer *grpc.Server) {
	pb.RegisterStoreServer(gRPCServer, &serverAPI{})
}

func (s *serverAPI) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	return &pb.SetResponse{Status: "Done", Key: req.GetKey()}, nil
}

func (s *serverAPI) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{Key: req.GetKey(), Value: "test"}, nil
}
func (s *serverAPI) Delete(ctx context.Context, req *pb.DelRequest) (*pb.DelResponse, error) {
	return &pb.DelResponse{Key: req.GetKey(), Status: "Done"}, nil
}

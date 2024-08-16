package server

import (
	"context"

	pb "github.com/PepsiKingIV/KeyValueDB/internal/server/gRPC"
	"google.golang.org/grpc"
)

type serverAPI struct {
	pb.UnimplementedStoreServer
}

func Register(gRPCServer *grpc.Server) {
	pb.RegisterStoreServer(gRPCServer, &serverAPI{})
}

func (s *serverAPI) Set(context.Context, *pb.SetRequest) (*pb.SetResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Get(context.Context, *pb.GetRequest) (*pb.GetResponse, error) {
	panic("implement me")
}
func (s *serverAPI) Delete(context.Context, *pb.DelRequest) (*pb.DelResponse, error) {
	panic("implement me")
}

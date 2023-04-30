package gapi

import (
	"context"

	"github.com/RomanIkonnikov93/keeper/server/internal/config"
	pb "github.com/RomanIkonnikov93/keeper/server/internal/proto"
	"github.com/RomanIkonnikov93/keeper/server/internal/repository"
	"github.com/RomanIkonnikov93/keeper/server/pkg/logging"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KeeperServiceServer struct {
	pb.UnimplementedKeeperServer
	rep    repository.Repository
	user   repository.IDRepository
	ping   repository.Pinger
	cfg    config.Config
	logger logging.Logger
}

func InitServices(rep repository.Reps, cfg config.Config, logger *logging.Logger) *KeeperServiceServer {

	return &KeeperServiceServer{
		rep:    rep.Rep,
		user:   rep.UserRep,
		ping:   rep.Ping,
		cfg:    cfg,
		logger: *logger,
	}
}

func (k *KeeperServiceServer) AddUser(ctx context.Context, in *pb.Auth) (*pb.Auth, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KeeperServiceServer) CheckUser(ctx context.Context, in *pb.Auth) (*pb.Auth, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KeeperServiceServer) Add(ctx context.Context, in *pb.Record) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KeeperServiceServer) Get(ctx context.Context, in *pb.Record) (*pb.Record, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KeeperServiceServer) GetAllByType(ctx context.Context, in *pb.Record) (*pb.Record, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KeeperServiceServer) UpdateByID(ctx context.Context, in *pb.Record) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KeeperServiceServer) DeleteByID(ctx context.Context, in *pb.Record) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

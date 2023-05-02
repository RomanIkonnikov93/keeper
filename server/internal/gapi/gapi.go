package gapi

import (
	"context"
	"errors"

	"github.com/RomanIkonnikov93/keeper/server/internal/authjwt"
	"github.com/RomanIkonnikov93/keeper/server/internal/config"
	"github.com/RomanIkonnikov93/keeper/server/internal/crypt"
	"github.com/RomanIkonnikov93/keeper/server/internal/models"
	pb "github.com/RomanIkonnikov93/keeper/server/internal/proto"
	"github.com/RomanIkonnikov93/keeper/server/internal/repository"
	"github.com/RomanIkonnikov93/keeper/server/pkg/logging"

	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KeeperServiceServer struct {
	pb.UnimplementedKeeperServer
	rep    repository.Repository
	user   repository.UsersRepository
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

func (k *KeeperServiceServer) RegistrationUser(ctx context.Context, in *pb.Auth) (*pb.Auth, error) {

	out := &pb.Auth{}

	if in.UserLogin == "" || in.UserPassword == "" {
		k.logger.Printf("%v", status.Error(codes.InvalidArgument, ""))
		return nil, status.Error(codes.InvalidArgument, "")
	}

	newID := ksuid.New().String()

	encrypted, err := crypt.Encrypt([]byte(in.UserPassword), []byte(k.cfg.SecretKey))
	if err != nil {
		k.logger.Error("")
		return nil, err
	}

	err = k.user.AddUser(ctx, newID, in.UserLogin, encrypted)
	if errors.Is(err, models.ErrConflict) {
		k.logger.Error(models.ErrConflict)
		return nil, status.Error(codes.AlreadyExists, "")
	} else if err != nil {
		k.logger.Error(err)
		return nil, err
	}

	token, err := authjwt.EncodeJWT(newID, k.cfg.JWTSecretKey)
	if err != nil {
		k.logger.Error(err)
		return nil, err
	}

	out.UserToken = token

	return out, nil
}

func (k *KeeperServiceServer) LoginUser(ctx context.Context, in *pb.Auth) (*pb.Auth, error) {

	out := &pb.Auth{}

	if in.UserLogin == "" || in.UserPassword == "" {
		k.logger.Printf("%v", status.Error(codes.InvalidArgument, ""))
		return nil, status.Error(codes.InvalidArgument, "")
	}

	encryptedPass, ID, err := k.user.CheckUser(ctx, in.UserLogin, in.UserPassword)
	if errors.Is(err, models.ErrNotExist) {
		k.logger.Error(models.ErrNotExist)
		return nil, status.Error(codes.NotFound, "")
	} else if err != nil {
		k.logger.Error(err)
		return nil, err
	}

	valid, err := crypt.CheckPasswordDecrypt(in.UserPassword, encryptedPass, []byte(k.cfg.SecretKey))
	if err != nil {
		k.logger.Error(err)
		return nil, err
	}
	if !valid {
		k.logger.Error(status.Error(codes.Unauthenticated, ""))
		return nil, status.Error(codes.Unauthenticated, "")
	}

	token, err := authjwt.EncodeJWT(ID, k.cfg.JWTSecretKey)
	if err != nil {
		k.logger.Error(err)
		return nil, err
	}

	out.UserToken = token

	return out, nil
}

func (k *KeeperServiceServer) AddRecord(ctx context.Context, in *pb.Record) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KeeperServiceServer) GetRecord(ctx context.Context, in *pb.Record) (*pb.Record, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KeeperServiceServer) GetAllRecordsByType(ctx context.Context, in *pb.Record) (*pb.List, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KeeperServiceServer) UpdateRecordByID(ctx context.Context, in *pb.Record) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KeeperServiceServer) DeleteRecordByID(ctx context.Context, in *pb.Record) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

package gapi

import (
	"context"
	"errors"
	"time"

	"github.com/RomanIkonnikov93/keeper/client/internal/config"
	"github.com/RomanIkonnikov93/keeper/client/internal/models"
	pb "github.com/RomanIkonnikov93/keeper/client/internal/proto"
	"github.com/RomanIkonnikov93/keeper/client/pkg/logging"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KeeperServiceClient struct {
	pb.KeeperClient
	Auth   models.Auth
	Record models.Record
	Store  models.Storage
	cfg    config.Config
	logger logging.Logger
}

func InitServices(cfg config.Config, logger *logging.Logger) (*KeeperServiceClient, error) {

	conn, err := grpc.Dial(cfg.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewKeeperClient(conn)

	return &KeeperServiceClient{
		KeeperClient: client,
		Auth: models.Auth{
			LastChanges: models.DefaultLastChangesTime,
		},
		Record: models.Record{},
		Store:  models.Storage{},
		cfg:    cfg,
		logger: *logger,
	}, nil
}

func (k *KeeperServiceClient) RegistrationUser() error {

	ctx := context.Background()

	in := &pb.Auth{
		UserLogin:    k.Auth.Login,
		UserPassword: k.Auth.Password,
	}

	res, err := k.KeeperClient.RegistrationUser(ctx, in)
	if err != nil {
		return err
	}

	k.Auth.Token = res.UserToken

	return nil
}

func (k *KeeperServiceClient) LoginUser() error {

	ctx := context.Background()

	in := &pb.Auth{
		UserLogin:    k.Auth.Login,
		UserPassword: k.Auth.Password,
	}

	res, err := k.KeeperClient.LoginUser(ctx, in)
	if err != nil {
		return err
	}

	k.Auth.Token = res.UserToken

	return nil
}

func (k *KeeperServiceClient) AddRecord(in *pb.Record) error {

	md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	switch in.RecordType {
	case models.Credentials:
	case models.Card:
	case models.File:

	}

	_, err := k.KeeperClient.AddRecord(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

//
//func (k *KeeperServiceClient) GetRecordByID(ctx context.Context, in *interface{}, opts ...grpc.CallOption) (*interface{}, error) {
//
//	md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
//	ctx := metadata.NewOutgoingContext(context.Background(), md)
//}
//
//func (k *KeeperServiceClient) GetAllRecordsByType(ctx context.Context, in *interface{}, opts ...grpc.CallOption) (*interface{}, error) {
//
//	md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
//	ctx := metadata.NewOutgoingContext(context.Background(), md)
//}
//
//func (k *KeeperServiceClient) UpdateRecordByID(ctx context.Context, in *interface{}, opts ...grpc.CallOption) (*emptypb.Empty, error) {
//
//	md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
//	ctx := metadata.NewOutgoingContext(context.Background(), md)
//}
//
//func (k *KeeperServiceClient) DeleteRecordByID(ctx context.Context, in *interface{}, opts ...grpc.CallOption) (*emptypb.Empty, error) {
//
//	md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
//	ctx := metadata.NewOutgoingContext(context.Background(), md)
//}

func (k *KeeperServiceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KeeperServiceClient) CheckChanges() error {

	md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	in := &pb.Record{
		RecordType: models.Credentials,
		CreatedAt:  k.Auth.LastChanges,
	}

	credentials, err := k.KeeperClient.CheckChanges(ctx, in)
	if errors.Is(err, status.Error(codes.NotFound, "")) {
	} else if err != nil {
		return err
	} else {
		for _, val := range credentials.GetNote() {
			k.Store.Credentials = append(k.Store.Credentials, models.Record{
				RecordID:    val.RecordID,
				RecordType:  val.RecordType,
				Description: val.Description,
				Metadata:    val.Metadata,
				Login:       val.Login,
				Password:    val.Password,
				CreatedAt:   val.CreatedAt,
			})

			last, err := time.Parse(models.DefaultLastChangesTime, k.Auth.LastChanges)
			if err != nil {
				return err
			}
			now, err := time.Parse(models.DefaultLastChangesTime, val.CreatedAt)
			if err != nil {
				return err
			}
			if now.After(last) {
				k.Auth.LastChanges = val.CreatedAt
			}
		}
	}

	in.RecordType = models.Card

	cards, err := k.KeeperClient.CheckChanges(ctx, in)
	if errors.Is(err, status.Error(codes.NotFound, "")) {
	} else if err != nil {
		return err
	} else {
		for _, val := range cards.GetNote() {
			k.Store.Cards = append(k.Store.Cards, models.Record{
				RecordID:    val.RecordID,
				RecordType:  val.RecordType,
				Description: val.Description,
				Metadata:    val.Metadata,
				Card:        val.Card,
				CreatedAt:   val.CreatedAt,
			})

			last, err := time.Parse(models.DefaultLastChangesTime, k.Auth.LastChanges)
			if err != nil {
				return err
			}
			now, err := time.Parse(models.DefaultLastChangesTime, val.CreatedAt)
			if err != nil {
				return err
			}
			if now.After(last) {
				k.Auth.LastChanges = val.CreatedAt
			}
		}
	}
	return nil
}

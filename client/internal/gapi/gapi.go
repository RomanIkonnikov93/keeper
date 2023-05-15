package gapi

import (
	"context"
	"errors"
	"sync"
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
)

type KeeperServiceClient struct {
	pb.KeeperClient
	Auth   models.Auth
	Record models.Record
	Store  models.Storage
	Mutex  *sync.RWMutex
	cfg    config.Config
	Logger logging.Logger
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
			LastChangesTime: models.DefaultLastChangesTime,
		},
		Record: models.Record{},
		Store: models.Storage{
			Credentials: make(map[int32]models.Record),
			Cards:       make(map[int32]models.Record),
		},
		Mutex:  &sync.RWMutex{},
		cfg:    cfg,
		Logger: *logger,
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
		k.Logger.Error(err)
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
		k.Logger.Error(err)
		return err
	}

	k.Auth.Token = res.UserToken

	return nil
}

func (k *KeeperServiceClient) AddRecord() error {

	in := k.SetRecordFields()

	md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	switch k.Record.RecordType {
	case models.Credentials:
	case models.Card:
	case models.File:

	}

	_, err := k.KeeperClient.AddRecord(ctx, in)
	if err != nil {
		k.Logger.Error(err)
		return err
	}

	if in.RecordType == models.Credentials || in.RecordType == models.Card {
		err = k.CheckChanges()
		if err != nil {
			k.Logger.Error(err)
			return err
		}
	}

	k.CleanRecordFields()

	return nil
}

func (k *KeeperServiceClient) GetRecordByID() error {

	switch k.Record.RecordType {
	case models.Credentials:

		k.Mutex.Lock()
		defer k.Mutex.Unlock()

		record, ok := k.Store.Credentials[k.Record.RecordID]
		if !ok {
			k.Logger.Error(models.ErrNotExist)
			return models.ErrNotExist
		}

		k.Record = record

	case models.Card:

		k.Mutex.Lock()
		defer k.Mutex.Unlock()

		record, ok := k.Store.Cards[k.Record.RecordID]
		if !ok {
			k.Logger.Error(models.ErrNotExist)
			return models.ErrNotExist
		}

		k.Record = record

	case models.File:
	}

	return nil
}

// func (k *KeeperServiceClient) GetAllRecordsByType(ctx context.Context, in *interface{}, opts ...grpc.CallOption) (*interface{}, error) {
//
//		md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
//		ctx := metadata.NewOutgoingContext(context.Background(), md)
//	}
func (k *KeeperServiceClient) UpdateRecordByID() error {

	in := k.SetRecordFields()

	md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	switch k.Record.RecordType {
	case models.Credentials:
	case models.Card:
	case models.File:

	}

	_, err := k.KeeperClient.UpdateRecordByID(ctx, in)
	if err != nil {
		k.Logger.Error(err)
		return err
	}

	if in.RecordType == models.Credentials || in.RecordType == models.Card {
		err = k.CheckChanges()
		if err != nil {
			k.Logger.Error(err)
			return err
		}
	}

	k.CleanRecordFields()

	return nil
}

func (k *KeeperServiceClient) DeleteRecordByID() error {

	in := k.SetRecordFields()

	md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err := k.KeeperClient.DeleteRecordByID(ctx, in)
	if err != nil {
		k.Logger.Error(err)
		return err
	}

	switch in.RecordType {
	case models.Credentials:

		k.Mutex.Lock()
		delete(k.Store.Credentials, in.RecordID)
		k.Mutex.Unlock()

	case models.Card:

		k.Mutex.Lock()
		delete(k.Store.Cards, in.RecordID)
		k.Mutex.Unlock()

	case models.File:

	}

	k.CleanRecordFields()

	return nil
}

func (k *KeeperServiceClient) Ping() error {
	//TODO implement me
	panic("implement me")
}

func (k *KeeperServiceClient) CheckChanges() error {

	md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	in := &pb.Record{
		RecordType: models.Credentials,
		CreatedAt:  k.Auth.LastChangesTime,
	}

	credentials, err := k.KeeperClient.CheckChanges(ctx, in)
	if errors.Is(err, status.Error(codes.NotFound, "")) {
	} else if err != nil {
		k.Logger.Error(err)
		return err
	} else {
		for _, val := range credentials.GetNote() {

			k.Mutex.Lock()
			k.Store.Credentials[val.RecordID] = models.Record{
				RecordID:    val.RecordID,
				RecordType:  val.RecordType,
				Description: val.Description,
				Metadata:    val.Metadata,
				Login:       val.Login,
				Password:    val.Password,
				CreatedAt:   val.CreatedAt,
			}
			k.Mutex.Unlock()

			last, err := time.Parse(models.DefaultLastChangesTime, k.Auth.LastChangesTime)
			if err != nil {
				k.Logger.Error(err)
				return err
			}
			now, err := time.Parse(models.DefaultLastChangesTime, val.CreatedAt)
			if err != nil {
				k.Logger.Error(err)
				return err
			}
			if now.After(last) {
				k.Auth.LastChangesTime = val.CreatedAt
			}
		}
	}

	in.RecordType = models.Card

	cards, err := k.KeeperClient.CheckChanges(ctx, in)
	if errors.Is(err, status.Error(codes.NotFound, "")) {
	} else if err != nil {
		k.Logger.Error(err)
		return err
	} else {
		for _, val := range cards.GetNote() {

			k.Mutex.Lock()
			k.Store.Cards[val.RecordID] = models.Record{
				RecordID:    val.RecordID,
				RecordType:  val.RecordType,
				Description: val.Description,
				Metadata:    val.Metadata,
				Card:        val.Card,
				CreatedAt:   val.CreatedAt,
			}
			k.Mutex.Unlock()

			last, err := time.Parse(models.DefaultLastChangesTime, k.Auth.LastChangesTime)
			if err != nil {
				k.Logger.Error(err)
				return err
			}
			now, err := time.Parse(models.DefaultLastChangesTime, val.CreatedAt)
			if err != nil {
				k.Logger.Error(err)
				return err
			}
			if now.After(last) {
				k.Auth.LastChangesTime = val.CreatedAt
			}
		}
	}
	return nil
}

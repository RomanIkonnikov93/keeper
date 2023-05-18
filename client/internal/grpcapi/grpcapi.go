package grpcapi

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/RomanIkonnikov93/keeper/client/internal/config"
	"github.com/RomanIkonnikov93/keeper/client/internal/models"
	pb "github.com/RomanIkonnikov93/keeper/client/internal/proto"
	"github.com/RomanIkonnikov93/keeper/client/pkg/logging"

	"github.com/golang/protobuf/ptypes/empty"
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
	Cfg    config.Config
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
			FileInfo:    make(map[int32]models.Record),
		},
		Mutex:  &sync.RWMutex{},
		Cfg:    cfg,
		Logger: *logger,
	}, nil
}

// RegistrationUser registers a new user and assigns a unique ID to him.
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

// LoginUser authorizes and identifies users.
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

// AddRecord adds a new record to the database.
func (k *KeeperServiceClient) AddRecord() error {

	in := k.SetRecordFields()

	md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	switch k.Record.RecordType {
	case models.File:
		path := filepath.Clean(k.Record.FilePath)
		file, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if len(file) > 1024*1024*2 { // 2Mb
			k.Logger.Error(models.ErrMaxFileSize)
			return models.ErrMaxFileSize
		}
		in.File = file

		fileName := filepath.Base(path)
		in.Metadata = fileName
	}

	_, err := k.KeeperClient.AddRecord(ctx, in)
	if err != nil {
		k.Logger.Error(err)
		return err
	}

	err = k.CheckChanges()
	if err != nil {
		k.Logger.Error(err)
		return err
	}

	k.CleanRecordFields()

	return nil
}

// GetRecordByID gets the record from the RAM storage (binary data is requested directly from the server).
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

		in := k.SetRecordFields()

		md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		out, err := k.KeeperClient.GetRecordByID(ctx, in)
		if err != nil {
			k.Logger.Error(err)
			return err
		}

		k.Record.Description = out.Description
		k.Record.Metadata = out.Metadata

		path := ""

		if k.Record.FilePath != "" {
			path = k.Record.FilePath + "/" + out.Metadata
		} else if k.Cfg.DownloadFilesPath != "" {
			path = k.Cfg.DownloadFilesPath + "/" + out.Metadata
		} else {
			dir, err := os.Getwd()
			if err != nil {
				k.Logger.Error(err)
				return err
			}
			path = dir + "/" + out.Metadata
		}

		err = os.WriteFile(filepath.Clean(path), out.File, 0666)
		if err != nil {
			k.Logger.Error(err)
			return err
		}

		return nil
	}

	return nil
}

// UpdateRecordByID updates the record in the database.
func (k *KeeperServiceClient) UpdateRecordByID() error {

	in := k.SetRecordFields()

	md := metadata.New(map[string]string{"usertoken": k.Auth.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	switch k.Record.RecordType {
	case models.File:
		if k.Record.FilePath != "" {
			path := filepath.Clean(k.Record.FilePath)
			file, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if len(file) > 1024*1024*2 { // 2Mb
				return models.ErrMaxFileSize
			}
			in.File = file

			fileName := filepath.Base(path)
			in.Metadata = fileName
		}
	}

	_, err := k.KeeperClient.UpdateRecordByID(ctx, in)
	if err != nil {
		k.Logger.Error(err)
		return err
	}

	err = k.CheckChanges()
	if err != nil {
		k.Logger.Error(err)
		return err
	}

	k.CleanRecordFields()

	return nil
}

// DeleteRecordByID changes the status of a record in the database to: deleted, the record is also deleted from RAM storage.
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

		k.Mutex.Lock()
		delete(k.Store.FileInfo, in.RecordID)
		k.Mutex.Unlock()

	}

	k.CleanRecordFields()

	return nil
}

// CheckChanges scans the database for new or changed records, depending on the last modification time (stored on the client).
// And duplicate data in RAM (excluding binary data).
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

	in.RecordType = models.File

	info, err := k.KeeperClient.CheckChanges(ctx, in)
	if errors.Is(err, status.Error(codes.NotFound, "")) {
	} else if err != nil {
		k.Logger.Error(err)
		return err
	} else {
		for _, val := range info.GetNote() {

			k.Mutex.Lock()
			k.Store.FileInfo[val.RecordID] = models.Record{
				RecordID:    val.RecordID,
				RecordType:  val.RecordType,
				Description: val.Description,
				Metadata:    val.Metadata,
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

// Ping checks the database connection.
func (k *KeeperServiceClient) Ping() error {
	in := &empty.Empty{}
	_, err := k.KeeperClient.Ping(context.Background(), in)
	if err != nil {
		return err
	}
	return nil
}

package gapi

import (
	"context"
	"errors"

	"github.com/RomanIkonnikov93/keeper/server/internal/authjwt"
	"github.com/RomanIkonnikov93/keeper/server/internal/cardvalid"
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

	out := &emptypb.Empty{}

	ID, err := authjwt.UserTokenValidation(ctx, k.cfg.JWTSecretKey)
	if err != nil {
		k.logger.Error(err)
		return nil, status.Error(codes.Unauthenticated, "")
	}

	in.UserID = ID

	switch in.RecordType {
	case models.Credentials:

		if in.Login == "" || in.Password == "" {
			k.logger.Error(status.Error(codes.InvalidArgument, ""))
			return nil, status.Error(codes.InvalidArgument, "")
		}

		encrypted, err := crypt.Encrypt([]byte(in.Password), []byte(k.cfg.SecretKey))
		if err != nil {
			k.logger.Error(err)
			return nil, err
		}

		in.Password = encrypted

	case models.Card:

		valid, err := cardvalid.CheckCard(in.Card)
		if !valid || err != nil {
			k.logger.Error(status.Error(codes.InvalidArgument, ""))
			return nil, status.Error(codes.InvalidArgument, "")
		}

		encrypted, err := crypt.Encrypt([]byte(in.Card), []byte(k.cfg.SecretKey))
		if err != nil {
			k.logger.Error(err)
			return nil, err
		}

		in.Card = encrypted

	case models.File:

		if len(in.File) == 0 {
			k.logger.Error(status.Error(codes.InvalidArgument, ""))
			return nil, status.Error(codes.InvalidArgument, "")
		}

	default:
		k.logger.Error(status.Error(codes.InvalidArgument, ""))
		return nil, status.Error(codes.InvalidArgument, "")
	}

	err = k.rep.Add(ctx, in)
	if err != nil {
		k.logger.Error(err)
		return nil, err
	}

	return out, nil
}

func (k *KeeperServiceServer) GetRecordByID(ctx context.Context, in *pb.Record) (*pb.Record, error) {

	out := &pb.Record{}

	ID, err := authjwt.UserTokenValidation(ctx, k.cfg.JWTSecretKey)
	if err != nil {
		k.logger.Error(err)
		return nil, status.Error(codes.Unauthenticated, "")
	}

	in.UserID = ID

	if in.RecordID == 0 {
		k.logger.Error(status.Error(codes.InvalidArgument, ""))
		return nil, status.Error(codes.InvalidArgument, "")
	}

	switch in.RecordType {
	case models.Credentials:

		out, err = k.rep.Get(ctx, in)
		if err != nil {
			k.logger.Error(err)
			return nil, err
		}

		password, err := crypt.Decrypt(out.Password, []byte(k.cfg.SecretKey))
		if err != nil {
			k.logger.Error(err)
			return nil, err
		}

		out.Password = string(password)

	case models.Card:

		out, err = k.rep.Get(ctx, in)
		if err != nil {
			k.logger.Error(err)
			return nil, err
		}

		card, err := crypt.Decrypt(out.Card, []byte(k.cfg.SecretKey))
		if err != nil {
			k.logger.Error(err)
			return nil, err
		}

		out.Card = string(card)

	case models.File:

		out, err = k.rep.Get(ctx, in)
		if err != nil {
			k.logger.Error(err)
			return nil, err
		}

	default:
		k.logger.Error(status.Error(codes.InvalidArgument, ""))
		return nil, status.Error(codes.InvalidArgument, "")
	}

	return out, nil
}

func (k *KeeperServiceServer) GetAllRecordsByType(ctx context.Context, in *pb.Record) (*pb.List, error) {

	out := &pb.List{}

	ID, err := authjwt.UserTokenValidation(ctx, k.cfg.JWTSecretKey)
	if err != nil {
		k.logger.Error(err)
		return nil, status.Error(codes.Unauthenticated, "")
	}

	in.UserID = ID

	switch in.RecordType {
	case models.Credentials:
	case models.Card:
	case models.File:
	default:
		k.logger.Error(status.Error(codes.InvalidArgument, ""))
		return nil, status.Error(codes.InvalidArgument, "")
	}

	data, err := k.rep.GetAllByType(ctx, in)
	if err != nil {
		k.logger.Error(err)
		return nil, err
	}

	for _, val := range data {
		switch val.RecordType {
		case models.Credentials:

			password, err := crypt.Decrypt(val.Password, []byte(k.cfg.SecretKey))
			if err != nil {
				k.logger.Error(err)
				return nil, err
			}

			out.Note = append(out.Note, &pb.Record{
				RecordID:    val.RecordID,
				RecordType:  val.RecordType,
				Description: val.Description,
				Metadata:    val.Metadata,
				Login:       val.Login,
				Password:    string(password),
				CreatedAt:   val.CreatedAt,
			})

		case models.Card:

			card, err := crypt.Decrypt(val.Card, []byte(k.cfg.SecretKey))
			if err != nil {
				k.logger.Error(err)
				return nil, err
			}

			out.Note = append(out.Note, &pb.Record{
				RecordID:    val.RecordID,
				RecordType:  val.RecordType,
				Description: val.Description,
				Metadata:    val.Metadata,
				Card:        string(card),
				CreatedAt:   val.CreatedAt,
			})

		case models.File:

			out.Note = append(out.Note, &pb.Record{
				RecordID:    val.RecordID,
				RecordType:  val.RecordType,
				Description: val.Description,
				Metadata:    val.Metadata,
				File:        val.File,
				CreatedAt:   val.CreatedAt,
			})

		default:
			k.logger.Error(status.Error(codes.InvalidArgument, ""))
			return nil, status.Error(codes.InvalidArgument, "")
		}
	}

	return out, nil
}

func (k *KeeperServiceServer) UpdateRecordByID(ctx context.Context, in *pb.Record) (*emptypb.Empty, error) {

	out := &emptypb.Empty{}

	ID, err := authjwt.UserTokenValidation(ctx, k.cfg.JWTSecretKey)
	if err != nil {
		k.logger.Error(err)
		return nil, status.Error(codes.Unauthenticated, "")
	}

	in.UserID = ID

	if in.RecordID == 0 {
		k.logger.Error(status.Error(codes.InvalidArgument, ""))
		return nil, status.Error(codes.InvalidArgument, "")
	}

	switch in.RecordType {
	case models.Credentials:

		if in.Login == "" || in.Password == "" {
			k.logger.Error(status.Error(codes.InvalidArgument, ""))
			return nil, status.Error(codes.InvalidArgument, "")
		}

		encrypted, err := crypt.Encrypt([]byte(in.Password), []byte(k.cfg.SecretKey))
		if err != nil {
			k.logger.Error(err)
			return nil, err
		}

		in.Password = encrypted

	case models.Card:

		valid, err := cardvalid.CheckCard(in.Card)
		if !valid || err != nil {
			k.logger.Error(status.Error(codes.InvalidArgument, ""))
			return nil, status.Error(codes.InvalidArgument, "")
		}

		encrypted, err := crypt.Encrypt([]byte(in.Card), []byte(k.cfg.SecretKey))
		if err != nil {
			k.logger.Error(err)
			return nil, err
		}

		in.Card = encrypted

	case models.File:

		if len(in.File) == 0 {
			k.logger.Error(status.Error(codes.InvalidArgument, ""))
			return nil, status.Error(codes.InvalidArgument, "")
		}

	default:
		k.logger.Error(status.Error(codes.InvalidArgument, ""))
		return nil, status.Error(codes.InvalidArgument, "")
	}

	err = k.rep.UpdateByID(ctx, in)
	if err != nil {
		k.logger.Error(err)
		return nil, err
	}

	return out, nil
}

func (k *KeeperServiceServer) DeleteRecordByID(ctx context.Context, in *pb.Record) (*emptypb.Empty, error) {

	out := &emptypb.Empty{}

	ID, err := authjwt.UserTokenValidation(ctx, k.cfg.JWTSecretKey)
	if err != nil {
		k.logger.Error(err)
		return nil, status.Error(codes.Unauthenticated, "")
	}

	in.UserID = ID

	if in.RecordID == 0 {
		k.logger.Error(status.Error(codes.InvalidArgument, ""))
		return nil, status.Error(codes.InvalidArgument, "")
	}

	switch in.RecordType {
	case models.Credentials:
	case models.Card:
	case models.File:
	default:
		k.logger.Error(status.Error(codes.InvalidArgument, ""))
		return nil, status.Error(codes.InvalidArgument, "")
	}

	err = k.rep.DeleteByID(ctx, in)
	if err != nil {
		k.logger.Error(err)
		return nil, err
	}

	return out, nil
}

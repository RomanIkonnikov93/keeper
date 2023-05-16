package gapi

import pb "github.com/RomanIkonnikov93/keeper/client/internal/proto"

// SetRecordFields sets input form fields before submitting to server.
func (k *KeeperServiceClient) SetRecordFields() *pb.Record {

	return &pb.Record{
		RecordID:    k.Record.RecordID,
		RecordType:  k.Record.RecordType,
		Description: k.Record.Description,
		Metadata:    k.Record.Metadata,
		Login:       k.Record.Login,
		Password:    k.Record.Password,
		Card:        k.Record.Card,
	}
}

// CleanRecordFields clears the entry form fields after submitting to server.
func (k *KeeperServiceClient) CleanRecordFields() {
	k.Record.RecordID = 0
	k.Record.RecordType = ""
	k.Record.Description = ""
	k.Record.Metadata = ""
	k.Record.Login = ""
	k.Record.Password = ""
	k.Record.Card = ""
	k.Record.File = nil
	k.Record.FilePath = ""
	k.Record.ActionType = ""
}

package repository

import (
	pb "github.com/RomanIkonnikov93/keeper/server/internal/proto"

	sq "github.com/Masterminds/squirrel"
)

// sqlBuilderForCredentials generates an update request to the database depending on the conditions.
func sqlBuilderForCredentials(in *pb.Record) (string, []interface{}, error) {

	n := "now()"

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := psql.Update("users_credentials").Set("created_at", n)
	if in.Description != "" {
		builder = builder.Set("description", in.Description)
	}
	if in.Metadata != "" {
		builder = builder.Set("metadata", in.Metadata)
	}
	if in.Login != "" {
		builder = builder.Set("user_login", in.Login)
	}
	if in.Password != "" {
		builder = builder.Set("user_password", in.Password)
	}

	builder = builder.Where(sq.Eq{"user_id": in.UserID, "record_id": in.RecordID, "del_flag": false})

	sql, args, err := builder.ToSql()
	if err != nil {
		return "", nil, err
	}
	return sql, args, nil
}

// sqlBuilderForCard generates an update request to the database depending on the conditions.
func sqlBuilderForCard(in *pb.Record) (string, []interface{}, error) {

	n := "now()"

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := psql.Update("users_cards").Set("created_at", n)
	if in.Description != "" {
		builder = builder.Set("description", in.Description)
	}
	if in.Metadata != "" {
		builder = builder.Set("metadata", in.Metadata)
	}
	if in.Card != "" {
		builder = builder.Set("user_card", in.Card)
	}

	builder = builder.Where(sq.Eq{"user_id": in.UserID, "record_id": in.RecordID, "del_flag": false})

	sql, args, err := builder.ToSql()
	if err != nil {
		return "", nil, err
	}
	return sql, args, nil
}

// sqlBuilderForFile generates an update request to the database depending on the conditions.
func sqlBuilderForFile(in *pb.Record) (string, []interface{}, error) {

	n := "now()"

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := psql.Update("users_files").Set("created_at", n)
	if in.Description != "" {
		builder = builder.Set("description", in.Description)
	}
	if in.Metadata != "" {
		builder = builder.Set("metadata", in.Metadata)
	}
	if in.File != nil {
		builder = builder.Set("user_file", in.File)
	}

	builder = builder.Where(sq.Eq{"user_id": in.UserID, "record_id": in.RecordID, "del_flag": false})

	sql, args, err := builder.ToSql()
	if err != nil {
		return "", nil, err
	}
	return sql, args, nil
}

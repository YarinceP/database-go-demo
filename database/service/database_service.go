package service

import (
	"context"
	"database/sql"
	"log"
)

type DatabaseService interface {
	QueryRowContext(ctx context.Context, query string, params interface{}) (interface{}, error)
}

type DatabaseServiceImpl struct {
	db *sql.DB
}

func NewDatabaseServiceImpl(db *sql.DB) *DatabaseServiceImpl {
	return &DatabaseServiceImpl{db: db}
}

func (receiver DatabaseServiceImpl) QueryRowContext(ctx context.Context, query string, params interface{}, repository interface{}) (interface{}, error) {

	row := receiver.db.QueryRowContext(ctx, query, params)

	err := row.Scan(repository)
	if err != nil {
		log.Println(err)
	}
	return repository, nil
}

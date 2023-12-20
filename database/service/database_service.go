package service

import (
	"context"
	"database/sql"
	"log"
	"reflect"
)

type DatabaseService interface {
	QueryRowContextReassign(ctx context.Context, query string, params interface{}) error
	QueryRowContextWithoutModel(ctx context.Context, query string, params interface{}) (interface{}, error)
	QueryRowContextWithReflectType(ctx context.Context, query string, params interface{}) (interface{}, error)
}

type DatabaseServiceImpl struct {
	db *sql.DB
}

func NewDatabaseServiceImpl(db *sql.DB) *DatabaseServiceImpl {
	return &DatabaseServiceImpl{db: db}
}

func (receiver DatabaseServiceImpl) QueryRowContext(ctx context.Context, query string, params interface{}, repository interface{}) error {
	row := receiver.db.QueryRowContext(ctx, query, params)
	err := row.Scan(repository)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (receiver DatabaseServiceImpl) QueryRowContextReassign(ctx context.Context, query string, params interface{}, repository interface{}) error {
	row := receiver.db.QueryRowContext(ctx, query, params)
	err := row.Scan(repository)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (receiver DatabaseServiceImpl) QueryRowContextWithoutModel(ctx context.Context, query string, params interface{}) (interface{}, error) {
	row := receiver.db.QueryRowContext(ctx, query, params)

	var result interface{}
	err := row.Scan(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (receiver DatabaseServiceImpl) QueryRowContextWithReflectType(ctx context.Context, query string, params interface{}, repositoryType reflect.Type) (interface{}, error) {
	row := receiver.db.QueryRowContext(ctx, query, params)

	repository := reflect.New(repositoryType).Interface()
	err := row.Scan(repository)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return repository, nil
}

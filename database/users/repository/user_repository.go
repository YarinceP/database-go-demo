package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/yarincep/database-go-demo/database/users/model"
	"github.com/yarincep/database-go-demo/database/users/queries"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userId int) (*model.User, error)
}

type UserRepositoryImplement struct {
	db *sql.DB
}

// NewUserRepository crea una nueva instancia de UserRepositoryImplement.
func NewUserRepository(db *sql.DB) *UserRepositoryImplement {
	return &UserRepositoryImplement{
		db: db,
	}
}

// GetUserByID obtiene un usuario por su ID desde la base de datos.
func (r *UserRepositoryImplement) GetUserByID(ctx context.Context, userID int) (*model.User, error) {
	row := r.db.QueryRowContext(ctx, queries.UserQueries.GetUserByID, userID)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}
	return user, nil
}

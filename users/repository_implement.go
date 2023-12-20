package users

import (
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImplement struct {
	db *sql.DB
}

func NewUserRepositoryImplement(db *sql.DB) *UserRepositoryImplement {
	return &UserRepositoryImplement{db: db}
}

func (r UserRepositoryImplement) Save(ctx context.Context, user *User) error {
	if r.db == nil {
		return errors.New("DB nil")
	}
	// Utilizamos prepared statements para prevenir inyección SQL
	_, err := r.db.ExecContext(ctx, UserQueries.InsertUser, &user.ID, &user.Name)
	if err != nil {
		return err
	}
	return nil
}

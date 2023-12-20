package service

import (
	"context"
	"fmt"
	"github.com/yarincep/database-go-demo/database/users/model"
	"github.com/yarincep/database-go-demo/database/users/repository"
	"time"
)

type UserServiceRepository interface {
	GetUserByID(userID int) (*model.User, error)
}

// UserService proporciona funciones de servicio para la entidad de usuario.
type UserService struct {
	userRepository repository.UserRepository
}

// NewUserService crea una nueva instancia de UserService.
func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// GetUserByID obtiene un usuario por su ID.
func (s *UserService) GetUserByID(userID int) (*model.User, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}
	return user, nil
}

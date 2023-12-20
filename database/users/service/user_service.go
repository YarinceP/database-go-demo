package service

import (
	"database-go-demo/database/users/model"
	"database-go-demo/database/users/repository"
	"fmt"
)

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
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}
	return user, nil
}

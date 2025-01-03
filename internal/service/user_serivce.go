package service

import (
	"errors"
	"go_mico_demo/internal/model"
	"go_mico_demo/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *model.User) error {
	existing, err := s.userRepo.GetByEmail(user.Email)
	if err == nil && existing != nil {
		return errors.New("email already exists")
	}
	return s.userRepo.Create(user)
}

func (s *UserService) GetUser(id uint) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

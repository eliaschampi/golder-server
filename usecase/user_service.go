package usecase

import (
	"context"
	"golden-server/domain/entity"
)

type userService struct {
	userRepository entity.UserRepository
}

func NewUserService(repo *entity.UserRepository) entity.UserService {
	return &userService{*repo}
}

func (s *userService) GetAll(c context.Context) ([]*entity.User, error) {
	return s.userRepository.GetAll(c)
}

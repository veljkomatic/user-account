package service

import (
	"context"
	"github.com/veljkomatic/user-account/pkg/user/domain"
	"github.com/veljkomatic/user-account/pkg/user/repository"
)

type UserService interface {
	GetUser(ctx context.Context, userID domain.UserID) (*domain.User, error)
	CreateUser(ctx context.Context, createUserParams *CreateUserParams, options ...UserOption) (*domain.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

var _ UserService = (*userService)(nil)

func (s *userService) GetUser(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	return s.repository.GetUserByID(ctx, userID)
}

func (s *userService) CreateUser(ctx context.Context, createUserParams *CreateUserParams, options ...UserOption) (*domain.User, error) {
	user := domain.NewUser(createUserParams.Name)

	for _, option := range options {
		option(user)
	}

	if err := s.repository.SaveUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

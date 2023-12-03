package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/veljkomatic/user-account/pkg/user/domain"
	"github.com/veljkomatic/user-account/pkg/user/service"
)

// UserService is a mock type for service.UserService
type UserService struct {
	mock.Mock
}

func (s *UserService) CreateUser(ctx context.Context, params *service.CreateUserParams, options ...service.UserOption) (*domain.User, error) {
	args := s.Called(ctx, params, options)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (s *UserService) GetUser(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	args := s.Called(ctx, userID)
	var user *domain.User
	if args.Get(0) != nil {
		user = args.Get(0).(*domain.User)
	}
	return user, args.Error(1)
}

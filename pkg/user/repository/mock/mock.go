package mock

import (
	"context"
	"github.com/stretchr/testify/mock"

	"github.com/veljkomatic/user-account/pkg/user/domain"
)

// UserRepository is a mock type for repository.UserRepository
type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *UserRepository) SaveUser(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

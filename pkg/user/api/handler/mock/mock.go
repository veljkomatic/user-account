package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/veljkomatic/user-account/pkg/user/api"
	"github.com/veljkomatic/user-account/pkg/user/api/handler"
)

// UserHandler is a mock type for handler.UserHandler
type UserHandler struct {
	mock.Mock
}

func (m *UserHandler) CreateUser(ctx api.APIContext, req *handler.CreateUserRequest) (*handler.CreateUserResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*handler.CreateUserResponse), args.Error(1)
}

func (m *UserHandler) GetUser(ctx api.APIContext, req *handler.GetUserRequest) (*handler.GetUserResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*handler.GetUserResponse), args.Error(1)
}

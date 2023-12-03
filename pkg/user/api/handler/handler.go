package handler

import (
	"github.com/veljkomatic/user-account/common/ptr"
	"github.com/veljkomatic/user-account/pkg/user/api"
	"github.com/veljkomatic/user-account/pkg/user/service"
)

type UserHandler interface {
	// CreateUser will create a new user and return the user object, currently it returns whole user object
	// it can be changed to return sanitized user object, using transformer
	CreateUser(
		ctx api.APIContext,
		request *CreateUserRequest,
	) (*CreateUserResponse, error)
	GetUser(
		ctx api.APIContext,
		request *GetUserRequest,
	) (*GetUserResponse, error)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

var _ UserHandler = (*userHandler)(nil)

func (u *userHandler) CreateUser(ctx api.APIContext, request *CreateUserRequest) (*CreateUserResponse, error) {
	createUserParams := service.CreateUserParams{
		Name: request.Name,
	}
	user, err := u.userService.CreateUser(ctx.Context(), ptr.From(createUserParams))
	if err != nil {
		return nil, err
	}
	return ptr.From(CreateUserResponse{
		User: user,
	}), nil
}

func (u *userHandler) GetUser(ctx api.APIContext, request *GetUserRequest) (*GetUserResponse, error) {
	user, err := u.userService.GetUser(ctx.Context(), request.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return ptr.From(GetUserResponse{
		User: user,
	}), nil
}

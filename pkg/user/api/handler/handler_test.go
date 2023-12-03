package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/veljkomatic/user-account/common/errors"
	"github.com/veljkomatic/user-account/pkg/user/api"
	"github.com/veljkomatic/user-account/pkg/user/domain"
	mockservice "github.com/veljkomatic/user-account/pkg/user/service/mock"
)

// Test for CreateUser
func TestCreateUser(t *testing.T) {
	mockService := new(mockservice.UserService)
	handler := NewUserHandler(mockService)

	ctx := api.NewAPIContext(context.Background())
	request := &CreateUserRequest{Name: "Test User"}

	mockService.On("CreateUser", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*service.CreateUserParams"), mock.Anything).Return(&domain.User{ID: "1", Name: "Test User"}, nil)

	response, err := handler.CreateUser(ctx, request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Test User", response.User.Name)

	mockService.AssertExpectations(t)
}

// Test for GetUser
func TestGetUser(t *testing.T) {
	mockService := new(mockservice.UserService)
	handler := NewUserHandler(mockService)

	ctx := api.NewAPIContext(context.Background())
	userID := "1" // Assuming domain.UserID is a type alias for string
	request := &GetUserRequest{UserID: domain.UserID(userID)}

	// Mock successful response
	mockService.On("GetUser", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("domain.UserID")).Return(&domain.User{ID: "1", Name: "Test User"}, nil)

	response, err := handler.GetUser(ctx, request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, domain.UserID(userID), response.User.ID)
	assert.Equal(t, "Test User", response.User.Name)

	// Clear or reset mock expectations before setting up the next scenario
	mockService.ExpectedCalls = nil
	mockService.Calls = nil

	// Mock error not found response
	errorUserID := domain.UserID("2")
	mockService.On("GetUser", mock.AnythingOfType("*context.emptyCtx"), errorUserID).Return(nil, errors.ErrNotFound)

	response, err = handler.GetUser(ctx, &GetUserRequest{UserID: errorUserID})

	assert.Error(t, err)
	assert.Nil(t, response)

	mockService.AssertExpectations(t)
}

package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	commonerrors "github.com/veljkomatic/user-account/common/errors"
	"github.com/veljkomatic/user-account/pkg/user/domain"
	mockrepository "github.com/veljkomatic/user-account/pkg/user/repository/mock"
)

// Test for GetUser in UserService
func TestUserService_GetUser(t *testing.T) {
	mockRepo := new(mockrepository.UserRepository)
	service := NewUserService(mockRepo)
	ctx := context.Background()
	testUserID := domain.UserID("test-id")

	// Test user found
	expectedUser := &domain.User{ID: testUserID, Name: "Test User"}
	mockRepo.On("GetUserByID", ctx, testUserID).Return(expectedUser, nil)
	user, err := service.GetUser(ctx, testUserID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)

	// Resetting mock
	mockRepo.ExpectedCalls = nil
	mockRepo.Calls = nil

	// Test user not found
	mockRepo.On("GetUserByID", ctx, testUserID).Return((*domain.User)(nil), commonerrors.ErrNotFound)
	user, err = service.GetUser(ctx, testUserID)
	assert.NoError(t, err)
	assert.Nil(t, user)

	// Resetting mock
	mockRepo.ExpectedCalls = nil
	mockRepo.Calls = nil

	// Test repository error
	mockRepo.On("GetUserByID", ctx, testUserID).Return((*domain.User)(nil), errors.New("db error"))
	user, err = service.GetUser(ctx, testUserID)
	assert.Error(t, err)
	assert.Nil(t, user)

	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(mockrepository.UserRepository)
	service := NewUserService(mockRepo)
	ctx := context.Background()

	// Test successful user creation
	createUserParams := &CreateUserParams{Name: "New User"}
	expectedUser := &domain.User{Name: createUserParams.Name}
	mockRepo.On("SaveUser", ctx, mock.AnythingOfType("*domain.User")).Return(nil)
	user, err := service.CreateUser(ctx, createUserParams)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.Name, user.Name)

	// Resetting mock
	mockRepo.ExpectedCalls = nil
	mockRepo.Calls = nil

	// Test error during user creation
	mockRepo.On("SaveUser", ctx, mock.AnythingOfType("*domain.User")).Return(errors.New("db error"))
	user, err = service.CreateUser(ctx, createUserParams)
	assert.Error(t, err)
	assert.Nil(t, user)

	mockRepo.AssertExpectations(t)
}

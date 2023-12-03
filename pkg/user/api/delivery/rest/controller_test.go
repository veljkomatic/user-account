package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/veljkomatic/user-account/pkg/user/api/handler"
	mockhandler "github.com/veljkomatic/user-account/pkg/user/api/handler/mock"
	"github.com/veljkomatic/user-account/pkg/user/domain"
)

func TestUserController_GetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Set up the controller and routes
	mockHandler := new(mockhandler.UserHandler)
	uc := newUserController(mockHandler)
	router.GET("/user/:userID", uc.GetUser)

	// Mock response setup for GetUser
	userID := uuid.New().String()
	mockResponse := &handler.GetUserResponse{User: &domain.User{ID: domain.NewUserID(userID), Name: "Test User"}}
	mockHandler.On("GetUser", mock.Anything, mock.AnythingOfType("*handler.GetUserRequest")).Return(mockResponse, nil)

	// Create request and recorder
	req, _ := http.NewRequest("GET", "/user/"+userID, nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions as needed to validate the response body, headers, etc.

	mockHandler.AssertExpectations(t)
}

func TestUserController_CreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Set up the controller and routes
	mockHandler := new(mockhandler.UserHandler)
	uc := newUserController(mockHandler)

	router.POST("/user", uc.CreateUser)

	// Mock response setup for CreateUser
	mockCreateUserReq := &handler.CreateUserRequest{Name: "Test User"}
	mockCreateUserResp := &handler.CreateUserResponse{User: &domain.User{ID: domain.NewUserID(uuid.New().String()), Name: "Test User"}}
	mockHandler.On("CreateUser", mock.Anything, mockCreateUserReq).Return(mockCreateUserResp, nil)

	// Create request and recorder
	marshaledReq, err := json.Marshal(mockCreateUserReq)
	assert.NoError(t, err)
	body := bytes.NewBufferString(string(marshaledReq))
	req, _ := http.NewRequest("POST", "/user", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	// Add more assertions as needed

	mockHandler.AssertExpectations(t)
}

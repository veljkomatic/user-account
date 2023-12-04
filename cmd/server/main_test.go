package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/veljkomatic/user-account/pkg/user/api/handler"
)

func TestMain(m *testing.M) {
	// Improve this code to use test database and test configuration running in separate docker container
	// currently it will use the same configuration as the server, which is not good for testing
	// also we need setup.sql and cleanup.sql scripts to setup and cleanup the test database
	// also add http-req.json that contains all test cases for the server
	go main()
	time.Sleep(3 * time.Second)
	m.Run()
}

func TestFunc(t *testing.T) {
	ctx := context.Background()

	serverURL := "http://localhost:8000/users"
	httpClient := http.Client{Timeout: time.Second * 3}

	createUserBody := handler.CreateUserRequest{Name: "Jane Doe"}
	createUserBodyJSON, err := json.Marshal(createUserBody)
	assert.NoError(t, err)
	reqCreateUser, err := http.NewRequestWithContext(ctx, http.MethodPost, serverURL, strings.NewReader(string(createUserBodyJSON)))
	assert.NoError(t, err)

	resUser, err := httpClient.Do(reqCreateUser)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resUser.StatusCode)
	resUserBody, err := io.ReadAll(resUser.Body)
	assert.NoError(t, err)

	var createUserRes handler.CreateUserResponse
	err = json.Unmarshal(resUserBody, &createUserRes)
	assert.NoError(t, err)
	user := createUserRes.User
	assert.Equal(t, "Jane Doe", user.Name)

	getUserURL := fmt.Sprintf("%s/%s", serverURL, user.ID)
	reqGetUser, err := http.NewRequestWithContext(ctx, http.MethodGet, getUserURL, nil)
	assert.NoError(t, err)

	resUser, err = httpClient.Do(reqGetUser)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resUser.StatusCode)
	resUserBody, err = io.ReadAll(resUser.Body)
	assert.NoError(t, err)

	var getUserRes handler.GetUserResponse
	err = json.Unmarshal(resUserBody, &getUserRes)
	assert.NoError(t, err)
	fetchedUser := getUserRes.User
	assert.Equal(t, user.ID, fetchedUser.ID)
	assert.Equal(t, user.Name, fetchedUser.Name)
}

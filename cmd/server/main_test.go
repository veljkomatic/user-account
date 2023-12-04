package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// Improve this code to use test database and test configuration running in separate docker container
	// currently it will use the same configuration as the server, which is not good for testing
	go main()
	time.Sleep(3 * time.Second)
	m.Run()
}

func TestFunc(t *testing.T) {
	ctx := context.Background()

	serverURL := "http://localhost:8000/users"

	createUserBody := strings.NewReader(`{"name":"Jane Doe"}`)
	reqCreateUser, err := http.NewRequestWithContext(ctx, http.MethodPost, serverURL, createUserBody)
	assert.NoError(t, err)

	httpClient := http.Client{Timeout: time.Second * 3}
	resUser, err := httpClient.Do(reqCreateUser)
	assert.NoError(t, err)
	_, err = io.ReadAll(resUser.Body)
	assert.NoError(t, err)
}

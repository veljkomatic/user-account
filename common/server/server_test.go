package server

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultIntServerConfig(t *testing.T) {
	cfg := NewDefaultIntServerConfig()
	assert.Equal(t, DefaultIntServerPort, cfg.Port)
	assert.Equal(t, defaultServerReadTimeout, cfg.ReadTimeout)
	assert.Equal(t, defaultServerWriteTimeout, cfg.WriteTimeout)
	assert.Equal(t, defaultServerIdleTimeout, cfg.IdleTimeout)
	assert.Equal(t, defaultServerShutdownTimeout, cfg.ShutdownTimeout)
}

func TestConfig_Validate(t *testing.T) {
	cfg := &Config{}
	err := cfg.Validate()
	assert.NoError(t, err)
	assert.Equal(t, defaultServerPort, cfg.Port)
	assert.Equal(t, defaultServerReadTimeout, cfg.ReadTimeout)
	assert.Equal(t, defaultServerWriteTimeout, cfg.WriteTimeout)
	assert.Equal(t, defaultServerIdleTimeout, cfg.IdleTimeout)
	assert.Equal(t, defaultServerShutdownTimeout, cfg.ShutdownTimeout)
}

func TestConfigureServer(t *testing.T) {
	cfg := &Config{
		Port:            8080,
		ReadTimeout:     100,
		WriteTimeout:    200,
		IdleTimeout:     300,
		ShutdownTimeout: 400,
	}
	server := ConfigureServer(cfg)
	assert.Equal(t, ":8080", server.srv.Addr)
	assert.Equal(t, 100*time.Second, server.srv.ReadTimeout)
	assert.Equal(t, 200*time.Second, server.srv.WriteTimeout)
	assert.Equal(t, 300*time.Second, server.srv.IdleTimeout)
	assert.Equal(t, 400*time.Second, server.shutdownTimeout)
}

// TestServer_ListenAndServeAndShutdown is a basic structure for testing.
// This test might require a more complex setup, possibly involving mocking or an integration test.
func TestServer_ListenAndServeAndShutdown(t *testing.T) {
	cfg := NewDefaultIntServerConfig()
	server := ConfigureServer(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start the server in a goroutine
	go func() {
		err := server.ListenAndServe(ctx, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		if !assert.ErrorIs(t, err, http.ErrServerClosed) {
			t.Errorf("expected server to close with http.ErrServerClosed, got %v", err)
		}
	}()

	// Give the server a moment to start
	time.Sleep(time.Second)

	// Make a request to the server to ensure it's running
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d", cfg.Port))
	if !assert.NoError(t, err) {
		t.Fatalf("failed to make request to the server: %v", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Shutdown the server
	err = server.Shutdown(context.Background())
	assert.NoError(t, err)
}

package mock

import (
	"time"

	"github.com/stretchr/testify/mock"
)

// RedisClient is a mock type for cache.RedisClient
type RedisClient struct {
	mock.Mock
}

func (m *RedisClient) Set(key string, val any, ttl time.Duration) error {
	args := m.Called(key, val, ttl)
	return args.Error(0)
}

func (m *RedisClient) Get(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func (m *RedisClient) GetByteArray(key string) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *RedisClient) Delete(keys ...string) (int64, error) {
	args := m.Called(keys)
	return args.Get(0).(int64), args.Error(1)
}

func (m *RedisClient) Close() error {
	args := m.Called()
	return args.Error(0)
}

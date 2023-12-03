package cache

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/veljkomatic/user-account/common/cache/redis/mock"
	"github.com/veljkomatic/user-account/pkg/user/domain"
)

func TestUserCache_Invalidate(t *testing.T) {
	mockRedis := new(mock.RedisClient)
	cache := NewUserCache(mockRedis)

	key := "userKey"
	mockRedis.On("Delete", []string{key}).Return(int64(1), nil)

	err := cache.Invalidate(context.Background(), key)
	assert.NoError(t, err)

	mockRedis.AssertExpectations(t)
}

func TestUserCache_Set(t *testing.T) {
	mockRedis := new(mock.RedisClient)
	cache := NewUserCache(mockRedis)

	key := "userKey"
	user := &domain.User{ /* initialize user data */ }
	userBytes, _ := json.Marshal(user)

	mockRedis.On("Set", key, userBytes, time.Duration(0)).Return(nil)

	err := cache.Set(context.Background(), key, user)
	assert.NoError(t, err)

	mockRedis.AssertExpectations(t)
}

func TestUserCache_Get(t *testing.T) {
	mockRedis := new(mock.RedisClient)
	cache := NewUserCache(mockRedis)

	key := "userKey"
	expectedUser := &domain.User{ /* initialize user data */ }
	userBytes, _ := json.Marshal(expectedUser)

	mockRedis.On("GetByteArray", key).Return(userBytes, nil)

	user, err := cache.Get(context.Background(), key)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	mockRedis.AssertExpectations(t)
}

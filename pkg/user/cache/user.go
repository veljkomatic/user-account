package cache

import (
	"context"
	"encoding/json"

	"github.com/veljkomatic/user-account/common/cache/redis"
	"github.com/veljkomatic/user-account/common/log"
	"github.com/veljkomatic/user-account/common/ptr"
	"github.com/veljkomatic/user-account/pkg/user/domain"
)

type UserCache interface {
	Invalidate(ctx context.Context, key string) error
	Set(ctx context.Context, key string, user *domain.User) error
	Get(ctx context.Context, key string) (*domain.User, error)
}

type userCache struct {
	redis *redis.Client
}

var _ UserCache = (*userCache)(nil)

func NewUserCache(redisClient *redis.Client) UserCache {
	return &userCache{
		redis: redisClient,
	}
}

func (c *userCache) Invalidate(ctx context.Context, key string) error {
	_, err := c.redis.Delete(key)
	if err != nil {
		log.Warn(
			ctx,
			"failed to invalidate user cache",
			log.Fields{
				"key":    key,
				"reason": err,
			},
		)
		return redis.HandleError(err)
	}
	return nil
}

func (c *userCache) Set(ctx context.Context, key string, user *domain.User) error {
	bytea, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = c.redis.Set(key, bytea, 0)
	if err != nil {
		log.Warn(
			ctx,
			"failed to set user cache",
			log.Fields{
				"reason": err,
				"key":    key,
				"value":  bytea,
			},
		)
		return redis.HandleError(err)
	}

	return nil
}

func (c *userCache) Get(ctx context.Context, key string) (*domain.User, error) {
	bytea, err := c.redis.GetByteArray(key)
	if err != nil {
		return nil, redis.HandleError(err)
	}
	var user domain.User
	err = json.Unmarshal(bytea, &user)
	if err != nil {
		return nil, err
	}
	return ptr.From(user), nil
}

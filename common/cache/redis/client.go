package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis"

	"github.com/veljkomatic/user-account/common/log"
)

// Client is a redis client interface
type Client interface {
	Set(key string, val any, ttl time.Duration) error
	Get(key string) (string, error)
	GetByteArray(key string) ([]byte, error)
	Delete(keys ...string) (int64, error)
	Close() error
}

type client struct {
	conn  redis.Cmdable
	close func() error
}

func MustConfigClient(ctx context.Context, cfg *Config) Client {
	var (
		conn    redis.Cmdable
		closeFn func() error
	)
	conn, closeFn = newBaseClient(cfg.URL, cfg.Password)
	if _, err := conn.Ping().Result(); err != nil {
		log.Error(ctx, err, "error configuring redis client", cfg)
		panic(err)
	}
	return &client{
		conn:  conn,
		close: closeFn,
	}
}

func newBaseClient(url, password string) (*redis.Client, func() error) {
	cli := redis.NewClient(
		&redis.Options{
			Addr:     url,
			Password: password,
		},
	)
	return cli, cli.Close
}

func (c *client) Set(key string, val any, ttl time.Duration) error {
	return c.conn.Set(key, val, ttl).Err()
}

func (c *client) Get(key string) (string, error) {
	return c.conn.Get(key).Result()
}

func (c *client) GetByteArray(key string) ([]byte, error) {
	return c.conn.Get(key).Bytes()
}

func (c *client) Delete(keys ...string) (int64, error) {
	return c.conn.Del(keys...).Result()
}

func (c *client) Close() error { return c.close() }

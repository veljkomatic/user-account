package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis"

	"github.com/veljkomatic/user-account/common/log"
)

type Client struct {
	conn  redis.Cmdable
	close func() error
}

func MustConfigClient(ctx context.Context, cfg *Config) *Client {
	var (
		conn    redis.Cmdable
		closeFn func() error
	)
	conn, closeFn = newBaseClient(cfg.URL, cfg.Password)
	if _, err := conn.Ping().Result(); err != nil {
		log.Error(ctx, err, "error configuring redis client", cfg)
		panic(err)
	}
	return &Client{
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

func (c *Client) Set(key string, val any, ttl time.Duration) error {
	return c.conn.Set(key, val, ttl).Err()
}

func (c *Client) Get(key string) (string, error) {
	return c.conn.Get(key).Result()
}

func (c *Client) GetByteArray(key string) ([]byte, error) {
	return c.conn.Get(key).Bytes()
}

func (c *Client) Delete(keys ...string) (int64, error) {
	return c.conn.Del(keys...).Result()
}

func (c *Client) Close() error { return c.close() }

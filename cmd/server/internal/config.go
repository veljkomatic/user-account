package internal

import (
	"github.com/veljkomatic/user-account/common/cache/redis"
	"github.com/veljkomatic/user-account/common/config"
	"github.com/veljkomatic/user-account/common/server"
	"github.com/veljkomatic/user-account/common/storage/sql/postgres"
)

// Config is the configuration for the server.
type Config struct {
	config.Config `yaml:",inline"`

	Server       *server.Config   `yaml:"server"`
	Redis        *redis.Config    `yaml:"redis"`
	Postgres     *postgres.Config `yaml:"postgres"`
	CacheEnabled bool             `yaml:"cache_enabled"`
}

func (cfg *Config) Validate() error {
	if err := cfg.Config.Validate(); err != nil {
		return err
	}
	if err := cfg.Server.Validate(); err != nil {
		return err
	}
	if err := cfg.Redis.Validate(); err != nil {
		return err
	}
	if err := cfg.Postgres.Validate(); err != nil {
		return err
	}
	return nil
}

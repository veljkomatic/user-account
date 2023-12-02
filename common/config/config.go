package config

import (
	"github.com/veljkomatic/user-account/common/log"
	"github.com/veljkomatic/user-account/common/server"
)

type Validator interface {
	Validate() error
}

type BaseConfig interface {
	GetIntServer() *server.Config
	GetLogger() *log.Config
}

var _ Validator = (*Config)(nil)

var _ BaseConfig = &Config{}

type Config struct {
	IntServer *server.Config `json:"int_server"            yaml:"int_server"`
	Logger    *log.Config    `json:"logger"                yaml:"logger"`
}

func (cfg *Config) GetIntServer() *server.Config {
	return cfg.IntServer
}

func (cfg *Config) GetLogger() *log.Config {
	return cfg.Logger
}

func (cfg *Config) Validate() error {
	if cfg.IntServer == nil {
		cfg.IntServer = server.NewDefaultIntServerConfig()
	}
	if cfg.IntServer.Port == 0 {
		cfg.IntServer.Port = server.DefaultIntServerPort
	}

	return cfg.IntServer.Validate()
}

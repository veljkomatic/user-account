package redis

import (
	"github.com/pkg/errors"

	"github.com/veljkomatic/user-account/common/log"
)

type Config struct {
	URL      string `json:"url"         yaml:"url"`
	Password string `json:"password"    yaml:"password"`
}

func (c *Config) LogFields() log.Fields {
	return log.Fields{
		"url":      c.URL,
		"password": c.Password,
	}
}

func (c *Config) Validate() error {
	if c.URL == "" {
		return errors.New("url field is missing")
	}
	return nil
}

package postgres

import (
	"github.com/pkg/errors"
)

type Config struct {
	URL      string `json:"url"                  yaml:"url"`
	Database string `json:"database"             yaml:"database"`
	User     string `json:"user"                 yaml:"user"`
	Password string `json:"password"             yaml:"password"`
}

func (cfg *Config) Validate() error {
	if cfg.URL == "" {
		return errors.New("missing postgres conn url")
	}
	if cfg.Database == "" {
		return errors.New("missing postgres database")
	}
	if cfg.User == "" {
		return errors.New("missing postgres user")
	}
	if cfg.Password == "" {
		return errors.New("missing postgres password")
	}

	return nil
}

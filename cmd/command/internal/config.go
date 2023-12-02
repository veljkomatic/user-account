package internal

import (
	"github.com/veljkomatic/user-account/common/storage/sql/postgres"
)

type CommandConfig struct {
	Postgres *postgres.Config `yaml:"postgres"`
}

func (cfg *CommandConfig) Validate() error {
	if err := cfg.Postgres.Validate(); err != nil {
		return err
	}
	return nil
}

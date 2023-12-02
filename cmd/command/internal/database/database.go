package database

import (
	"github.com/urfave/cli/v2"
	"github.com/veljkomatic/user-account/cmd/command/internal"
	"github.com/veljkomatic/user-account/common/environment"
)

func NewCmd(config *internal.CommandConfig, envAccessor environment.EnvAccessor) *cli.Command {
	return &cli.Command{
		Name:  "database",
		Usage: "Used for working with database",
		Subcommands: []*cli.Command{
			NewMigrateUpCmd(config, envAccessor),
		},
	}
}

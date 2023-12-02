package database

import (
	"os"

	"github.com/fatih/color"
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/veljkomatic/user-account/cmd/command/internal"
	"github.com/veljkomatic/user-account/common/environment"
	"github.com/veljkomatic/user-account/common/storage/sql/postgres/migration"
)

const (
	path = "file://migrations"
)

func NewMigrateUpCmd(config *internal.CommandConfig, envAccessor environment.EnvAccessor) *cli.Command {
	return &cli.Command{
		Name: "migrate-up",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "db", Usage: "db migrate up", Value: "default"},
		},
		Usage: "Used for database migration",
		Action: func(c *cli.Context) error {
			if !environment.Is(envAccessor, environment.EnvDevelopment) {
				color.Red("This can only be run in development environment")
				os.Exit(1)
			}

			cfg := config.Postgres
			color.White("Executing migrations from: %s", color.BlueString(path))
			color.White("Database URL: %s", color.BlueString(cfg.URL))

			mig := migration.MustConfigure(cfg, path)
			defer mig.Close()

			err := mig.Up()
			if errors.Is(err, migrate.ErrNoChange) {
				color.Green("all migrations executed")
				return nil
			}
			if err != nil {
				return errors.Wrap(err, "migrate up")
			}

			color.Green("migrations successfully ran")

			return nil
		},
	}
}

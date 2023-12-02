package migration

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/veljkomatic/user-account/common/log"
	"github.com/veljkomatic/user-account/common/storage/sql/postgres"
)

func MustConfigure(cfg *postgres.Config, path string) *migrate.Migrate {
	mig, err := Configure(cfg, path)
	if err != nil {
		log.Error(context.TODO(), err, "could not configure postgres")

		os.Exit(1)
	}

	return mig
}

func Configure(cfg *postgres.Config, path string) (*migrate.Migrate, error) {
	db, err := DBConnect(cfg)

	_, err = db.Exec("SELECT 1")
	if err != nil {
		return nil, errors.Wrap(err, "check connection")
	}

	driver, err := pg.WithInstance(db, &pg.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "init driver")
	}

	m, err := migrate.NewWithDatabaseInstance(
		path, cfg.Database, driver,
	)
	if err != nil {
		return nil, errors.Wrap(err, "init migrate")
	}

	return m, nil
}

func DBConnect(cfg *postgres.Config) (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable",
			cfg.User, cfg.Password, cfg.URL, cfg.Database,
		),
	)
	if err != nil {
		return nil, errors.Wrap(err, "open DB connection")
	}
	return db, nil
}

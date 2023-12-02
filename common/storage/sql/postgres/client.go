package postgres

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/veljkomatic/user-account/common/log"
)

var (
	ctxDefaultTimeout = 60 * time.Second
	ctxOverrideOnce   = sync.Once{}
)

var clientDefaults = options{
	WithReadTimeout(30 * time.Second),  // Driver default is 10
	WithWriteTimeout(30 * time.Second), // Driver default is 5
	WithDialTimeout(5 * time.Second),   // Driver's default.
	WithApplicationName(""),
}

type IDB = bun.IDB

// TransactionFunc is the transaction function being run by the client.
type TransactionFunc = func(ctx context.Context, tx bun.Tx) error

// clientConfig is the bun postgres client configuration structure
type clientConfig struct {
	pgCfg *Config

	readTimeout  time.Duration
	writeTimeout time.Duration
	dialTimeout  time.Duration

	appName string
}

var (
	_ bun.IDB = (*Client)(nil)
	_ IDB     = (*Client)(nil)
)

// Client is the Bun's DB object wrapper and postgresql client.
// Client should always be constructed either by the MustInitClient function or with NewClient
// Using Client without calling the constructor first can lead to nil pointers and/or undefined behaviour
type Client struct {
	DB *bun.DB

	cfg *clientConfig
}

// NewClient parses the config and returns a new *Client
// It does not validate the configuration as it will already be validated in the time called.
func NewClient(ctx context.Context, cfg *Config, opts ...Option) (*Client, error) {
	cli := &Client{cfg: &clientConfig{pgCfg: cfg}}
	if err := clientDefaults.applyDriverOpts(cli); err != nil {
		return nil, errors.Wrap(err, "failed applying default driver option")
	}
	if err := (options(opts)).applyDriverOpts(cli); err != nil {
		return nil, errors.Wrap(err, "failed applying driver option")
	}

	appName := cli.cfg.appName

	connector := pgdriver.NewConnector(
		pgdriver.WithAddr(cli.cfg.pgCfg.URL),
		pgdriver.WithUser(cli.cfg.pgCfg.User),
		pgdriver.WithPassword(cli.cfg.pgCfg.Password),
		pgdriver.WithDatabase(cli.cfg.pgCfg.Database),
		pgdriver.WithReadTimeout(cli.cfg.readTimeout),
		pgdriver.WithWriteTimeout(cli.cfg.writeTimeout),
		pgdriver.WithDialTimeout(cli.cfg.dialTimeout),
		pgdriver.WithApplicationName(appName),
		pgdriver.WithInsecure(true), // Add this line to disable SSL
	)
	db := bun.NewDB(sql.OpenDB(connector), pgdialect.New(), bun.WithDiscardUnknownColumns())

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "failed creating bun client: ")
	}

	cli.DB = db
	if err := (clientDefaults).applyNonDriverOpts(cli); err != nil {
		return nil, errors.Wrap(err, "failed applying default non driver option")
	}
	if err := (options(opts)).applyNonDriverOpts(cli); err != nil {
		return nil, errors.Wrap(err, "failed applying non driver option")
	}

	return cli, nil
}

// MustInitClient will try to initialize the client with the passed config, environment and options.
// If it fails it will not return an error but fatal.
func MustInitClient(ctx context.Context, cfg *Config, opts ...Option) *Client {
	cli, err := NewClient(ctx, cfg, opts...)
	if err != nil {
		log.Error(ctx, err, "error configuring Postgres2 client")
		panic(err)
	}

	return cli
}

// Close closes the database and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server
// to finish.
//
// It is rare to Close a DB, as the DB handle is meant to be
// long-lived and shared between many goroutines.
func (c *Client) Close() error { return c.DB.Close() }

// RunInTx runs the function in a transaction. If the function returns an error,
// the transaction is rolled back. Otherwise, the transaction is committed.
// If passed opts are nil it will create an empty opts object and not panic, so you are free to do so.
// If nil opts are passed directly to bun's RunInTx, it will cause a nil pointer error.
func (c *Client) RunInTx(ctx context.Context, opts *sql.TxOptions, fn TransactionFunc) error {
	if opts == nil {
		opts = &sql.TxOptions{}
	}

	return c.DB.RunInTx(ctx, opts, fn)
}

// NewInsertModel returns a new insert query with the passed model as the model parameter.
func (c *Client) NewInsertModel(model any) *bun.InsertQuery {
	return c.DB.NewInsert().Model(model)
}

// NewSelectModel returns a new select query with the passed model as the model parameter.
func (c *Client) NewSelectModel(model any) *bun.SelectQuery {
	return c.DB.NewSelect().Model(model)
}

// NewUpdateModel returns a new update query with the passed model as the model parameter.
func (c *Client) NewUpdateModel(model any) *bun.UpdateQuery {
	return c.DB.NewUpdate().Model(model)
}

// NewDeleteModel returns a new delete query with the passed model as the model parameter.
func (c *Client) NewDeleteModel(model any) *bun.DeleteQuery {
	return c.DB.NewDelete().Model(model)
}

// NewInsert is a wrapper around's bun NewInsert()
func (c *Client) NewInsert() *bun.InsertQuery { return c.DB.NewInsert() }

// NewSelect is a wrapper around's bun NewSelect()
func (c *Client) NewSelect() *bun.SelectQuery { return c.DB.NewSelect() }

// NewUpdate is a wrapper around's bun NewUpdate()
func (c *Client) NewUpdate() *bun.UpdateQuery { return c.DB.NewUpdate() }

// NewDelete is a wrapper around's bun NewDelete()
func (c *Client) NewDelete() *bun.DeleteQuery { return c.DB.NewDelete() }

// NewRaw is a wrapper around's bun NewRaw()
func (c *Client) NewRaw(query string, args ...interface{}) *bun.RawQuery {
	return c.DB.NewRaw(query, args...)
}

// BeginTx is a wrapper around's bun BeginTx()
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (bun.Tx, error) {
	return c.DB.BeginTx(ctx, opts)
}

// Option allows configuring the postgres2 client
// Option is a struct not a func because some options have to be applied before the driver is initialized
// and some options need the databse object.
// To unify the exposed API, and allow functional options for drivers we use this wrapper around functional options.
// Internally it's a little bit more complicated but allows us more configuration with functional options.
type Option struct {
	opt    func(c *Client) error
	driver bool
}

// Apply applies the option
func (o Option) apply(c *Client) error { return o.opt(c) }

type options []Option

// applyDriverOpts will apply all driver options.
// This should be called before the db is initialized and set or it will return an error.
// These options include network and timeout configuration.
func (oo options) applyDriverOpts(c *Client) error {
	if c.DB != nil {
		return errors.New("driver options must be applied before the db is initialized and set")
	}
	for _, o := range oo.driverOpts() {
		if err := o.apply(c); err != nil {
			return err
		}
	}
	return nil
}

// applyNonDriverOpts will apply all non driver options.
// This should be called after the db is initialized and set or it will return an error.
// These options include global context timeout and query hook configurations.
func (oo options) applyNonDriverOpts(c *Client) error {
	if c.DB == nil {
		return errors.New("non driver options must be applied after the db is initialized and set")
	}
	for _, o := range oo.nonDriverOpts() {
		if err := o.apply(c); err != nil {
			return err
		}
	}
	return nil
}

// driverOpts returns all options with the driver flag set to true
func (oo options) driverOpts() options {
	driverOpts := make(options, 0)
	for _, o := range oo {
		if o.driver {
			driverOpts = append(driverOpts, o)
		}
	}
	return driverOpts
}

// nonDriverOpts returns all options with the driver flag set to false
func (oo options) nonDriverOpts() options {
	driverOpts := make(options, 0)
	for _, o := range oo {
		if !o.driver {
			driverOpts = append(driverOpts, o)
		}
	}
	return driverOpts
}

// WithReadTimeout will set the client's deafult read timeout.
// This read timeout will be used if no context with a timeout or deadline is passed.
func WithReadTimeout(dur time.Duration) Option {
	return Option{
		opt: func(c *Client) error {
			c.cfg.readTimeout = dur
			return nil
		},
		driver: true,
	}
}

// WithWriteTimeout will set the client's default write timeout.
// This write timeout will be used if no context with a timeout or deadline is passed.
func WithWriteTimeout(dur time.Duration) Option {
	return Option{
		opt: func(c *Client) error {
			c.cfg.writeTimeout = dur
			return nil
		},
		driver: true,
	}
}

// WithDialTimeout will set the client's default dial timeout.
func WithDialTimeout(dur time.Duration) Option {
	return Option{
		opt: func(c *Client) error {
			c.cfg.dialTimeout = dur
			return nil
		},
		driver: true,
	}
}

// WithApplicationName sets the application name in the driver options.
func WithApplicationName(name string) Option {
	return Option{
		opt: func(c *Client) error {
			c.cfg.appName = name
			return nil
		},
		driver: true,
	}
}

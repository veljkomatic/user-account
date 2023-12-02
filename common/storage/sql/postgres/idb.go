package postgres

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

func (c *Client) Dialect() schema.Dialect {
	return c.DB.Dialect()
}

func (c *Client) NewValues(model interface{}) *bun.ValuesQuery {
	return c.DB.NewValues(model)
}

func (c *Client) NewCreateTable() *bun.CreateTableQuery {
	return c.DB.NewCreateTable()
}

func (c *Client) NewDropTable() *bun.DropTableQuery {
	return c.DB.NewDropTable()
}

func (c *Client) NewCreateIndex() *bun.CreateIndexQuery {
	return c.DB.NewCreateIndex()
}

func (c *Client) NewDropIndex() *bun.DropIndexQuery {
	return c.DB.NewDropIndex()
}

func (c *Client) NewTruncateTable() *bun.TruncateTableQuery {
	return c.DB.NewTruncateTable()
}

func (c *Client) NewAddColumn() *bun.AddColumnQuery {
	return c.DB.NewAddColumn()
}

func (c *Client) NewDropColumn() *bun.DropColumnQuery {
	return c.DB.NewDropColumn()
}

func (c *Client) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return c.DB.QueryContext(ctx, query, args...)
}

func (c *Client) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return c.DB.ExecContext(ctx, query, args...)
}

func (c *Client) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return c.DB.QueryRowContext(ctx, query, args...)
}

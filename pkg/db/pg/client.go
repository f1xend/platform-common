package pg

import (
	"context"
	"github.com/f1xend/platform-common/pkg/db"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type pgClient struct {
	masterDBC db.DB
}

func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to coonect to db: %v", err)
	}
	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}

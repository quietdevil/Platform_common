package pg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quietdevil/Platform_common/pkg/db"
)

type DBClient struct {
	db db.DB
}

func NewDBClient(ctx context.Context, dsn string) (db.Client, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &DBClient{
		db: &PgPool{dbc: pool},
	}, nil

}

func (c *DBClient) Close() error {
	c.db.Close()
	return nil
}

func (c *DBClient) DB() db.DB {
	return c.db
}

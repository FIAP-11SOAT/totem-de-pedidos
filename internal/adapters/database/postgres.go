package dbadapter

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

type DatabaseAdapter struct {
	client *pgx.Conn
}

func New() *DatabaseAdapter {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	connStr := fmt.Sprintf("%s://%s:%s@%s/%s%s",
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		url.QueryEscape(os.Getenv("DB_PASS")),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_OPTIONS"),
	)

	client, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		panic(err)
	}

	if err = client.Ping(ctx); err != nil {
		slog.Error("error ping database", err.Error(), err)
		panic(err)
	}

	return &DatabaseAdapter{client: client}
}

func (d *DatabaseAdapter) Client() *pgx.Conn {
	return d.client
}

func (d *DatabaseAdapter) DataBaseHeatlh(ctx context.Context) error {
	return d.client.Ping(ctx)
}

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

// TODO: implementar pool
// TODO: implementar abstração para transaction

type DatabaseAdapter struct {
	Client *pgx.Conn
}

type Input struct {
	DBDrive   string
	DBUser    string
	DBPass    string
	DBHost    string
	DBName    string
	DBOptions string
}

func New(input Input) *DatabaseAdapter {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	connStr := fmt.Sprintf("%s://%s:%s@%s/%s%s",
		input.DBDrive,
		input.DBUser,
		url.QueryEscape(input.DBPass),
		input.DBHost,
		input.DBName,
		input.DBOptions,
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

	return &DatabaseAdapter{Client: client}
}

func (d *DatabaseAdapter) DataBaseHeatlh(ctx context.Context) error {
	return d.Client.Ping(ctx)
}

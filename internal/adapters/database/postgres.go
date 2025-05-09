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
	Client *pgx.Conn
}

type Input struct {
	Db_driver  string
	Db_user    string
	Db_pass    string
	Db_host    string
	Db_name    string
	Db_options string
}

func New(input Input) *DatabaseAdapter {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	connStr := fmt.Sprintf("%s://%s:%s@%s/%s%s",
		input.Db_driver,
		input.Db_user,
		url.QueryEscape(input.Db_pass),
		input.Db_host,
		input.Db_name,
		input.Db_options,
	)

	fmt.Println("********************", connStr)

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

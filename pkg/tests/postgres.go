// package tests

// import (
// 	"context"
// 	"log"
// 	"path/filepath"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/testcontainers/testcontainers-go"
// 	"github.com/testcontainers/testcontainers-go/modules/postgres"
// 	"github.com/testcontainers/testcontainers-go/wait"
// )

// func CreatePostgresDataBase(t *testing.T) string {
// 	ctx := context.Background()

// 	dbName := "totempedidos"
// 	dbUser := "totempedidos"
// 	dbPassword := "totempedidos"

// 	postgresContainer, err := postgres.Run(ctx,
// 		"postgres:16-alpine",
// 		postgres.WithInitScripts(filepath.Join("migrations", "0001_init.sql")),
// 		postgres.WithDatabase(dbName),
// 		postgres.WithUsername(dbUser),
// 		postgres.WithPassword(dbPassword),
// 		testcontainers.WithWaitStrategy(
// 			wait.ForLog("database system is ready to accept connections").
// 				WithOccurrence(2).
// 				WithStartupTimeout(5*time.Second)),
// 	)
// 	assert.NoError(t, err)
// 	defer func() {
// 		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
// 			log.Printf("failed to terminate container: %s", err)
// 		}
// 	}()

// 	dbURL, err := postgresContainer.ConnectionString(ctx)
// 	assert.NoError(t, err)
// 	return dbURL
// }

package tests

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	log "github.com/sirupsen/logrus"
)

var db *sql.DB

func CreatePostgresDataBase(m *testing.T) string {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=totempedidos",
			"POSTGRES_USER=totempedidos",
			"POSTGRES_DB=totempedidos",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://totempedidos:totempedidos@%s/totempedidos?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseUrl)

	// resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// defer func() {
	// 	if err := pool.Purge(resource); err != nil {
	// 		log.Fatalf("Could not purge resource: %s", err)
	// 	}
	// }()

	return databaseUrl
}

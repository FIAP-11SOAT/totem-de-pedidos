package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"

	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/api"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseAdapter := dbadapter.New(dbadapter.Input{
		Db_driver:  os.Getenv("DB_DRIVER"),
		Db_user:    os.Getenv("DB_USER"),
		Db_pass:    os.Getenv("DB_PASS"),
		Db_host:    os.Getenv("DB_HOST"),
		Db_name:    os.Getenv("DB_NAME"),
		Db_options: os.Getenv("DB_OPTIONS"),
	})

	app := echo.New()
	app.Use(middleware.CORS())
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())
	api.Routers(app, databaseAdapter)

	app.Logger.Info(app.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}

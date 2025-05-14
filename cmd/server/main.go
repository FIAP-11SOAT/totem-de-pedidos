package main

import (
	"os"

	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/api"
	routes "github.com/FIAP-11SOAT/totem-de-pedidos/internal/api/routers"
	"github.com/labstack/echo/v4"
)

func main() {

	databaseAdapter := dbadapter.New(dbadapter.Input{
		Db_driver:  os.Getenv("DB_DRIVER"),
		Db_user:    os.Getenv("DB_USER"),
		Db_pass:    os.Getenv("DB_PASS"),
		Db_host:    os.Getenv("DB_HOST"),
		Db_name:    os.Getenv("DB_NAME"),
		Db_options: os.Getenv("DB_OPTIONS"),
	})

	echoEngine := echo.New()

	routes := routes.New(echoEngine)
	serverApi := api.New(routes.SetupRouters(databaseAdapter), os.Getenv("SERVER_PORT"))
	serverApi.ListenAndServe()
}

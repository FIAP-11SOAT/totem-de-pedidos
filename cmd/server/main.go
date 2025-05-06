package main

import (
	"os"

	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/api"
	routes "github.com/FIAP-11SOAT/totem-de-pedidos/internal/api/routers"
	"github.com/labstack/echo/v4"
)

// main inicializa o adaptador de banco de dados, configura o servidor web e inicia a aplicação na porta definida pela variável de ambiente SERVER_PORT.
func main() {
	databaseAdapter := dbadapter.New()
	
	echoEngine := echo.New()
	
	routes := routes.New(echoEngine)
	serverApi := api.New(routes.SetupRouters(databaseAdapter), os.Getenv("SERVER_PORT"))
	serverApi.ListenAndServe()
}

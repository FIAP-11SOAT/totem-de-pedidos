package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type api struct {
	echoEngine *echo.Echo
	port       string
}

// New cria e retorna uma nova instância da estrutura api com o Echo engine e a porta especificados.
func New(echoEngine *echo.Echo, port string) *api {
	return &api{echoEngine: echoEngine, port: port}
}

func (a *api) ListenAndServe() {
	a.echoEngine.Use(middleware.Logger())
	a.echoEngine.Use(middleware.Recover())

	err := a.echoEngine.Start(":" + a.port)
	if err != nil {
		fmt.Println(err)
		return
	}
}

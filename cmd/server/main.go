package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/api"
)

func getEnvOrDefault(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func main() {

	isProduction := getEnvOrDefault("PROFILE", "dev") == "prod"
	if !isProduction {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	databaseAdapter := dbadapter.New(dbadapter.Input{
		DBDrive:   os.Getenv("DB_DRIVER"),
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASS"),
		DBHost:    os.Getenv("DB_HOST"),
		DBName:    os.Getenv("DB_NAME"),
		DBOptions: os.Getenv("DB_OPTIONS"),
	})

	app := echo.New()
	app.Use(middleware.CORS())
	app.Use(middleware.Recover())
	//app.Use(middleware.Logger())
	api.Routers(app, databaseAdapter)
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

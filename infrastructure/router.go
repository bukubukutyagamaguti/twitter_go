package infrastructure

import (
	"api/server/interfaces/controllers"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	// Echo instance
	e := echo.New()
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Route "/"
	e.POST("/login", func(c echo.Context) error { return twitterController.Login(c) })

	// Route "/api"
	api := e.Group("/api")

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

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

	// controllers
	twitterController := controllers.NewTwitterController(NewSqlHandler(), NewTokenHandler())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Route "/"
	e.POST("/login", func(c echo.Context) error { return twitterController.Login(c) })

	// Route "/api"
	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}))
	api.GET("/home", func(c echo.Context) error { return twitterController.RelatedPost(c) })
	api.POST("/post", func(c echo.Context) error { return twitterController.CreatePost(c) })
	api.POST("/follow/:id", func(c echo.Context) error { return twitterController.CreateFollow(c) })
	api.POST("/refollow/:id", func(c echo.Context) error { return twitterController.DeleteFollow(c) })

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

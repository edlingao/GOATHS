package main

import (
	"embed"
	"exercise-app/routes"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed all:public
var public embed.FS

func main() {
	e := echo.New()
  e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
    Root: "public",
    Browse: false,
    HTML5: false,
    Filesystem: http.FS(public),
  }))

  e.Use(routes.GetUserIntoContext)

	routes.INDEX.Build(e)
  routes.AUTH.Build(e)
  routes.EXERCISE.Build(e)

	log.Fatal(e.Start(os.Getenv("GO_PORT")))
}

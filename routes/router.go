package routes

import "github.com/labstack/echo/v4"


type Router interface {
  Build( c *echo.Echo )
}


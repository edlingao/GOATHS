package routes

import (
	"exercise-app/controllers"
	"exercise-app/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthRoutes string

const AUTH AuthRoutes = "/auth"

func Guard(e *echo.Group) {
	e.Use(auth)
}

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := c.Request().Cookie("token")

		if err != nil || !models.VerifyAPIKey(token.Value) {
			log.Print("TOKEN ", err)
			return c.Redirect(http.StatusTemporaryRedirect, "/auth/signin")
		}

		return next(c)
	}
}

func GetUserIntoContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token, err := c.Request().Cookie("token")
		if err != nil {
			return next(c)
		}

		username, err := models.GetUsernameFromToken(token.Value)
		if err != nil {
			return next(c)
		}

		user, err := models.USER.GetByUsername(username)
		if err != nil {
			return next(c)
		}

		c.Set("user", user)
		return next(c)
	}
}

func (a AuthRoutes) Build(e *echo.Echo) {
	auth_routes := e.Group(string(a))
	auth_controller := controllers.AuthController{}
	auth_routes.GET("/register", auth_controller.RegisterView)
	auth_routes.POST("/register", auth_controller.RegisterUser)

	auth_routes.GET("/signin", auth_controller.LoginView)
	auth_routes.POST("/signin", auth_controller.SignIn)

  auth_routes.GET("/logout", auth_controller.SignOut)
}

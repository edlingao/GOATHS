package controllers

import (
	"exercise-app/models"
	"exercise-app/utils"
	view_auth "exercise-app/view/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct{}
type RegisterBody struct {
	Username string
	Email    string
	Password string
}

type SigninStatus struct {
	Success bool
	Message string
}

func (a AuthController) RegisterView(c echo.Context) error {
	return utils.Render(
		c,
		view_auth.Register(view_auth.RegisterViewModel{
      Title: "Register",
    }),
		http.StatusOK,
	)
}

func (a AuthController) LoginView(c echo.Context) error {
	return utils.Render(
		c,
		view_auth.SignIn(view_auth.SignInViewModel{
      Title: "Sign In",
    }),
		http.StatusOK,
	)
}

func (a AuthController) RegisterUser(c echo.Context) error {
	email := c.FormValue("email")
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := models.USER.Register(email, username, password)

	if err != nil {
		return utils.Render(
      c,
      view_auth.RegisterForm(view_auth.RegisterViewModel{
        Title: "Register",
        Error: err,
        InputErrors: map[string]string{
          "password": err.Error(),
        },
      }),
      http.StatusOK,
    )
  }

	signed_api, err := models.GenerateAPIKey(user.Username)

	if err != nil {
		return utils.Render(
      c,
      view_auth.RegisterForm(view_auth.RegisterViewModel{
        Title: "Register",
        Error: err,
        InputErrors: map[string]string{
          "password": "Error generating API key",
        },
      }),
      http.StatusOK,
    )
	}

	return setCookieAndRedirect(c, signed_api)
}

func (a AuthController) SignIn(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := models.USER.Login(email, password)

	if err != nil {
		return utils.Render(
      c,
      view_auth.SignInForm(view_auth.SignInViewModel{
        Title: "Sign In",
        Error: err,
        InputErrors: map[string]string{
          "password": "Invalid email or password",
        },
      }),
      http.StatusOK,
    )
  }

	signed_api, err := models.GenerateAPIKey(user.Username)

	if err != nil {
		return c.JSON(
			http.StatusBadGateway,
			SigninStatus{Success: false, Message: "Error generating API key"},
		)
	}

	return setCookieAndRedirect(c, signed_api)
}

func (a AuthController) SignOut(c echo.Context) error {
	c.SetCookie(models.RemoveCookie())
	c.Response().Header().Set("HX-Location", "/auth/signin")
	return c.JSON(http.StatusOK, SigninStatus{Success: true, Message: "Success"})
}

func setCookieAndRedirect(c echo.Context, signed_api string) error {
	cookie := models.GetCookie(signed_api)
	c.SetCookie(cookie)
	c.Response().Header().Set("HX-Location", "/")
	return c.JSON(http.StatusOK, SigninStatus{Success: true, Message: "Success"})
}

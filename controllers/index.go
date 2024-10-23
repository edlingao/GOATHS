package controllers

import (
	"errors"
	"exercise-app/components"
	"exercise-app/models"
	"exercise-app/utils"
	view_index "exercise-app/view/index"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndexController struct{}

var indexModel models.Index

func (i IndexController) Show(c echo.Context) error {
  user_req := c.Get("user")
  user, ok := user_req.(models.User)

  var err error
  if !ok {
    err = errors.New("User not found")
  }
  if !ok {
    return c.Redirect(http.StatusTemporaryRedirect, "/auth/signin")
  }
	return utils.Render(
		c,
		view_index.Index(
      view_index.IndexViewModel{
        Title: "Index",
        Error: err,
        Authenticated: ok,
        User: user,
      },
		),
		http.StatusOK,
  )
}

func (i IndexController) ChangeName(c echo.Context) error {
	newName := c.Param("name")

	indexModel.SetName(newName)

	return utils.Render(
		c,
		view_index.Index(
      view_index.IndexViewModel{
        Title: "Index",
        Authenticated: false,
      },
		),
		http.StatusOK,
	)
}

func (i IndexController) AddCount(c echo.Context) error {
	count := 1

	indexModel.AddCount(count)

	return utils.Render(c,
		components.Count(indexModel.GetCount()),
		http.StatusOK,
	)
}

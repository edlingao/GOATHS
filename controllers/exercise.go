package controllers

import (
	"exercise-app/components"
	"exercise-app/models"
	"exercise-app/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ExerciseController struct{}

func (e ExerciseController) Show(c echo.Context) error {
	user_req := c.Get("user")
	user, ok := user_req.(models.User)

	if !ok {
		return c.Redirect(http.StatusTemporaryRedirect, "/auth/signin")
	}

	excercises, err := models.EXERCISE.GetByUserID(user.ID)

	if err != nil {
		return c.HTML(
			http.StatusBadGateway,
			err.Error(),
		)
	}
	return utils.Render(
		c,
		components.ExerciseTable(
			components.ExerciseTableViewModel{
				Exercises: excercises,
			},
		),
		http.StatusOK,
	)
}

func (e ExerciseController) Create(c echo.Context) error {
	user_req := c.Get("user")
	user, ok := user_req.(models.User)

	if !ok {
		return c.Redirect(http.StatusTemporaryRedirect, "/auth/signin")
	}

	date := time.Now()
	feeling := c.FormValue("feeling")
  hours := c.FormValue("hours")
  minutes := c.FormValue("minutes")
	calories := c.FormValue("calories")
	calories_int, _ := strconv.ParseInt(calories, 10, 64)
  hours_int, _ := strconv.ParseInt(hours, 10, 64)
  minutes_int, _ := strconv.ParseInt(minutes, 10, 64)

	exercise := models.Exercise{
		User_ID:  user.ID,
		Date:     date,
		Feeling:  feeling,
    Hours: int(hours_int),
    Minutes: int(minutes_int),
		Calories: int(calories_int),
	}

	_, err := models.EXERCISE.Create(exercise)

	if err != nil {
		return utils.Render(
			c,
			components.ExerciseForm(
				components.ExerciseFormViewModel{
					Error: map[string]string{
						"general": err.Error(),
					}},
			),
			http.StatusOK,
		)
	}
	_, error := models.EXERCISE.GetByUserID(user.ID)

	if error != nil {
		return utils.Render(
			c,
			components.ExerciseForm(
				components.ExerciseFormViewModel{
					Error: map[string]string{
						"general": err.Error(),
					}},
			),
			http.StatusOK,
		)
	}

  c.Response().Header().Set("HX-Trigger", "{ \"exercises:loaded\": \"\" }")

  return utils.Render(
    c,
    components.ExerciseForm(
      components.ExerciseFormViewModel{},
    ),
    http.StatusOK,
  )
}

func (e ExerciseController) Delete(c echo.Context) error {
  id := c.Param("id")
  id_int, _ := strconv.ParseInt(id, 10, 64)

  _, err := models.EXERCISE.Delete(int(id_int))

  if err != nil {
    return c.HTML(
      http.StatusBadGateway,
      err.Error(),
    )
  }
  
  c.Response().Header().Set("HX-Trigger", "{ \"exercises:loaded\": \"\" }")

  return c.JSON(http.StatusOK, map[string]string{
    "message": "Exercise deleted",
  })
}


package routes

import (
	"exercise-app/controllers"

	"github.com/labstack/echo/v4"
)

type ExerciseRoutes string

const EXERCISE ExerciseRoutes = "/exercise"

func (ex ExerciseRoutes) Build(e *echo.Echo) {
	exercise_routes := e.Group(string(ex))
	exercise_controller := controllers.ExerciseController{}
	exercise_routes.GET("", exercise_controller.Show)
  exercise_routes.POST("/new", exercise_controller.Create)
  exercise_routes.DELETE("/delete/:id", exercise_controller.Delete)
}

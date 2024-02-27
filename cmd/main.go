package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mikhailovv/myexercise/handler"
)

func main() {
	app := echo.New()

	authHandler := handler.AuthHandler{}
	exercisesHandler := handler.ExercisesHandler{}

	authMiddleware := middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "cookie:token",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == "123", nil
		},
	})

	app.POST("/cookie", authHandler.CookieHandler)
	app.POST("/login", authHandler.LoginHandler)

	exerciseGroup := app.Group("/exercises", authMiddleware)
	exerciseGroup.GET("", exercisesHandler.GetExercisesHandler)
	exerciseGroup.POST("", exercisesHandler.CreateExercisesHandler)

	trainingGroup := app.Group("/trainings", authMiddleware)
	trainingGroup.POST("", handler.AddTraining)
	trainingGroup.GET("/:id", handler.GetTrainingByID)
	trainingGroup.POST("/:id/exercises/:exercise_id", handler.AddExerciseToTraining)
	trainingGroup.POST("/:id/exercises/:exercise_id/sets", handler.AddSetsToExercise)
	trainingGroup.POST("/:id/completed", handler.AddSetsToExercise)

	app.Start(":3000")
}

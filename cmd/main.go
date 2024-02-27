package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mikhailovv/myexercise/handler"
)

func main() {
	app := echo.New()

	authHandler := handler.AuthHandler{}
	exercisesHandler := handler.ExercisesHandler{}
	// trainingsHandler := handler.TrainingsHandler{}

	app.POST("/cookie", authHandler.CookieHandler)
	app.POST("/login", authHandler.LoginHandler)

	app.GET("/exercises", exercisesHandler.GetExercisesHandler)
	app.POST("/exercises", exercisesHandler.CreateExercisesHandler)

	app.POST("/trainings", handler.AddTraining)
	app.GET("/trainings/:id", handler.GetTrainingByID)
	app.POST("/trainings/:id/exercises/:exercise_id", handler.AddExerciseToTraining)
	app.POST("/trainings/:id/exercises/:exercise_id/sets", handler.AddSetsToExercise)
	app.POST("/trainings/:id/completed", handler.AddSetsToExercise)

	app.Start(":3000")
}

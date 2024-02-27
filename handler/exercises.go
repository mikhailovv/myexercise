package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikhailovv/myexercise/model"
)

type ExercisesHandler struct {
}

func (e ExercisesHandler) GetExercisesHandler(c echo.Context) error {
	exerciseRepository := model.ExerciseRepository{}
	exercises, error := exerciseRepository.GetExercises()
	if error != nil {
		c.JSON(http.StatusBadRequest, "Problem with fetch data")
	}

	return c.JSON(http.StatusOK, exercises)
}

func (e ExercisesHandler) CreateExercisesHandler(c echo.Context) error {
	exercise := new(model.Exercise)
	if err := c.Bind(exercise); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	exerciseRepository := model.ExerciseRepository{}
	savedExercise, saveError := exerciseRepository.Save(*exercise)
	if saveError != nil {
		return c.JSON(http.StatusCreated, "can't save exercise")
	}

	fmt.Println("Created new exercise", exercise.Name, " with ID ", exercise.ID)
	return c.JSON(http.StatusCreated, savedExercise)
}

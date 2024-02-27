package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikhailovv/myexercise/model"
)

// Set struct to represent a set within an exercise
type Set struct {
	Weight  int `json: "weight"`
	Repeats int `json: "repeats"`
}

// Exercise struct to store information about an exercise
type Exercise struct {
	Id   int    `json:"exercise_id"`
	Name string `json:"name"`
	Sets []Set  `json:"sets"`
}

// Training struct to store information about a training session
type Training struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	CreatedAt   time.Time        `json:"created_at"`
	CompletedAt time.Time        `json:"completed_at"`
	Exercises   map[int]Exercise `json:"exercises"`
	IsCompleted bool             `json:"is_completed"`
}

// Example of how to create a Training instance
var trainings = map[int]Training{
	1: {
		ID:          1,
		Name:        "Strength Training",
		CreatedAt:   time.Now(),
		CompletedAt: time.Now().Add(time.Hour * 2), // Assuming it finishes 2 hours after creation
		Exercises: map[int]Exercise{
			2: {
				Id:   2,
				Name: "Bench",
				Sets: []Set{
					{Weight: 100, Repeats: 10},
					{Weight: 110, Repeats: 8},
					{Weight: 120, Repeats: 6},
				},
			},
			3: {
				Id:   3,
				Name: "Barbel",
				Sets: []Set{
					{Weight: 80, Repeats: 8},
					{Weight: 90, Repeats: 6},
				},
			},
		},
	},
}

func GetTrainingByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid training ID"})
	}

	training, ok := trainings[id]
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Training not found"})
	}

	return c.JSON(http.StatusOK, training)
}

func AddExerciseToTraining(c echo.Context) error {
	// Get the training ID from the URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid training ID"})
	}

	// Check if the training exists
	training, ok := trainings[id]
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Training not found"})
	}
	if training.IsCompleted {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "You can't edit a completed training"})
	}

	exerciseId, err := strconv.Atoi(c.Param("exercise_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid exercise ID"})
	}

	// Find the index of the exercise in the training's exercises list
	if training.Exercises[exerciseId].Sets != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Exercise with ID=%d exists", exerciseId)})
	}

	exerciseRepository := model.ExerciseRepository{}
	exercise := exerciseRepository.Get(exerciseId)

	if exercise.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Exercise with ID=%d not exists", exerciseId)})
	}
	training.Exercises[exerciseId] = Exercise{Id: exerciseId, Name: exercise.Name}

	return c.JSON(http.StatusOK, training)
}

func AddSetsToExercise(c echo.Context) error {
	var sets []Set
	if err := c.Bind(&sets); err != nil {
		return err
	}

	// Get the training ID from the URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid training ID"})
	}

	// Check if the training exists
	training, ok := trainings[id]
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Training not found"})
	}
	if training.IsCompleted {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "You can't edit a completed training"})
	}

	exerciseId, err := strconv.Atoi(c.Param("exercise_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid exercise ID"})
	}

	exercise, ok := training.Exercises[exerciseId]
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Exercise not found in training"})
	}

	exercise.Sets = sets

	// Update the exercise in the training
	training.Exercises[exerciseId] = exercise

	return c.JSON(http.StatusCreated, exercise)
}

func generateUniqueID(trainings map[int]Training) int {
	maxID := 0
	for id := range trainings {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}

func AddTraining(c echo.Context) error {
	// Parse request body to get the details of the new training
	var newTraining Training
	if err := c.Bind(&newTraining); err != nil {
		return err
	}

	// Generate a unique ID for the new training
	newTrainingID := generateUniqueID(trainings)

	// Set the ID and CreatedAt fields for the new training
	newTraining.ID = newTrainingID
	newTraining.CreatedAt = time.Now()
	newTraining.Exercises = map[int]Exercise{}

	// Add the new training to the trainings map
	trainings[newTrainingID] = newTraining

	// Return the newly created training
	return c.JSON(http.StatusCreated, newTraining)
}

func CompleteTraining(c echo.Context) error {
	// Get the training ID from the URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid training ID"})
	}

	// Check if the training exists
	training, ok := trainings[id]
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Training not found"})
	}

	training.CompletedAt = time.Now()
	training.IsCompleted = true

	return c.JSON(http.StatusOK, training)
}

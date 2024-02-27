package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestAddSetsToExercise(t *testing.T) {
	// Create a new request with a JSON body containing sets
	requestBody := []byte(`[
        {"weight": 100, "repeats": 10},
        {"weight": 110, "repeats": 8}
    ]`)
	req, err := http.NewRequest("POST", "/trainings/1/exercises/2/sets", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create an Echo instance and set up the router
	e := echo.New()
	e.POST("/trainings/1/exercises/2/sets", AddSetsToExercise)

	// Perform the request
	e.ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	// Check the response body
	expectedResponseBody := `{"weight":100,"repeats":10},{"weight":110,"repeats":8}`
	if rr.Body.String() != expectedResponseBody {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedResponseBody)
	}
}

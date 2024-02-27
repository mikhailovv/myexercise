package model

type Exercise struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ExerciseRepository struct {
}

var exercisesStore = []Exercise{
	{ID: 1, Name: "Жим штанги"},
	{ID: 2, Name: "Жим гантелей сидя"},
	{ID: 3, Name: "Тяга нижнего блока"},
	{ID: 4, Name: "Жим ногами"},
	{ID: 5, Name: "Брусья"},
}

func (er ExerciseRepository) GetExercises() (exercises []Exercise, error error) {
	return exercisesStore, nil
}

func (er ExerciseRepository) Save(exercise Exercise) (Exercise, error) {
	lastExerciseId := exercisesStore[len(exercisesStore)-1].ID
	exercise.ID = lastExerciseId + 1

	exercisesStore = append(exercisesStore, exercise)

	return exercise, nil
}

func (er ExerciseRepository) Get(ID int) Exercise {
	var foundExercise Exercise
	for _, exercise := range exercisesStore {
		if exercise.ID == ID {
			foundExercise = exercise
			break
		}
	}

	return foundExercise
}

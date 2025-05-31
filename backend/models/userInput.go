package models

type UserInput struct {
    Age           int    `json:"age"`
    Height        int    `json:"height"`
    Weight        int    `json:"weight"`
    Gender        string `json:"gender"`
    ActivityLevel string `json:"activity_level"`
    Goal          string `json:"goal"`
    ExerciseMin   int    `json:"exercise_minutes"`
}
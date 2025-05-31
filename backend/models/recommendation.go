package models

type Recommendation struct {
    CaloriesNeeded      int    `json:"calories_needed"`
    MenuRecommendation  string `json:"menu_recommendation"`
    SportRecommendation string `json:"sport_recommendation"`
}
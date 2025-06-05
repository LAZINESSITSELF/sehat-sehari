package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Macro struct {
    ProteinG float64 `bson:"protein_g" json:"protein_g"`
    FatG     float64 `bson:"fat_g" json:"fat_g"`
    CarbsG   float64 `bson:"carbs_g" json:"carbs_g"`
}

type MenuItem struct {
    Name     string  `bson:"name" json:"name"`
    Calories float64 `bson:"calories" json:"calories"`
    Macros   Macro   `bson:"macros" json:"macros"`
}

type ActivityItem struct {
    Name           string  `bson:"name" json:"name"`
    DurationMin    int     `bson:"duration_min" json:"duration_min"`
    CaloriesBurned float64 `bson:"calories_burned" json:"calories_burned"`
}

type Recommendation struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID       primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
    Date         primitive.DateTime `bson:"date" json:"date"`
    CalorieNeed  float64            `bson:"calorie_need" json:"calorie_need"`
    Macros       Macro              `bson:"macros" json:"macros"`
    Menu         []MenuItem         `bson:"menu" json:"menu"`
    Activities   []ActivityItem     `bson:"activities" json:"activities"`
}

package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
    ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID             primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
    Age                int                `bson:"age" json:"age"`
    Gender             string             `bson:"gender" json:"gender"`
    HeightCm           float64            `bson:"height_cm" json:"height_cm"`
    WeightKg           float64            `bson:"weight_kg" json:"weight_kg"`
    ActivityLevel      string             `bson:"activity_level" json:"activity_level"`
    Goal               string             `bson:"goal" json:"goal"`
    ExerciseDuration   int                `bson:"exercise_duration" json:"exercise_duration"`
    UpdatedAt          primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}

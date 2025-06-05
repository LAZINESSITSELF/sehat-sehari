package services

import (
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"backend/models"
)

var MLClient *resty.Client

func InitMLClient() {
	MLClient = resty.New().
		SetHostURL(os.Getenv("ML_SERVICE_URL")).
		SetTimeout(10 * time.Second).
		SetHeader("Content-Type", "application/json")
}

// MLRequest sesuai dengan schema input di ML Service (FastAPI)
type MLRequest struct {
	Age              int     `json:"age"`
	Gender           string  `json:"gender"`
	HeightCm         float64 `json:"height_cm"`
	WeightKg         float64 `json:"weight_kg"`
	ActivityLevel    string  `json:"activity_level"`
	Goal             string  `json:"goal"`
	ExerciseDuration int     `json:"exercise_duration_min"`
}

// MLResponse sesuai dengan schema response dari ML Service
type MLResponse struct {
	CalorieNeed float64              `json:"calorie_need"`
	Macros      models.Macro         `json:"macros"`
	Menu        []models.MenuItem    `json:"menu"`
	Activities  []models.ActivityItem `json:"activities"`
}

// CallMLService memanggil endpoint ML Service dan mengembalikan hasil deserialisasi MLResponse
func CallMLService(req MLRequest) (*MLResponse, error) {
	resp, err := MLClient.R().
		SetBody(req).
		SetResult(&MLResponse{}).
		Post("")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*MLResponse), nil
}

package services

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "os"

    "backend/models"
)

// CallMLService mengirim input ke ML microservice dan mengembalikan rekomendasi.
func CallMLService(input models.UserInput) (models.Recommendation, error) {
    var reco models.Recommendation

    jsonData, err := json.Marshal(input)
    if err != nil {
        return reco, err
    }

    mlServiceURL := os.Getenv("ML_SERVICE_URL")
    if mlServiceURL == "" {
        return reco, errors.New("ML_SERVICE_URL tidak diset")
    }

    resp, err := http.Post(mlServiceURL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return reco, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return reco, errors.New("ML service error: " + resp.Status)
    }

    err = json.NewDecoder(resp.Body).Decode(&reco)
    return reco, err
}
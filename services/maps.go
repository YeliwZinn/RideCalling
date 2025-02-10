package services

import (
    "encoding/json"
    "fmt"
    //"math"
    "net/http"
    "uber-clone/config"
    //"uber-clone/db"
    //"go.mongodb.org/mongo-driver/bson"
)

// DistanceMatrixResponse is the structure of the response from the Distance Matrix API
type DistanceMatrixResponse struct {
    Rows []struct {
        Elements []struct {
            Distance struct {
                Value float64 `json:"value"` // Distance in meters
            } `json:"distance"`
            Duration struct {
                Value float64 `json:"value"` // Duration in seconds
            } `json:"duration"`
        } `json:"elements"`
    } `json:"rows"`
    Status string `json:"status"` // Add status check
}

// GetDistance calculates the distance and duration between two locations, and also calculates the fare
func GetDistance(originLat, originLng, destLat, destLng float64, vehicleType string) (float64, float64, float64, error) {
    apiKey := config.MustGetEnv("DISTANCEMATRIXAI_API_KEY")
    url := fmt.Sprintf(
        "https://api.distancematrix.ai/maps/api/distancematrix/json?origins=%f,%f&destinations=%f,%f&key=%s",
        originLat, originLng, destLat, destLng, apiKey,
    )

    resp, err := http.Get(url)
    if err != nil {
        return 0, 0, 0, fmt.Errorf("API request failed: %v", err)
    }
    defer resp.Body.Close()

    var data DistanceMatrixResponse
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return 0, 0, 0, fmt.Errorf("failed to parse response: %v", err)
    }

    // Check API status
    if data.Status != "OK" {
        return 0, 0, 0, fmt.Errorf("API error: %s", data.Status)
    }

    if len(data.Rows) == 0 || len(data.Rows[0].Elements) == 0 {
        return 0, 0, 0, fmt.Errorf("no distance or duration data found")
    }

    // Convert meters to kilometers and seconds to minutes
    distance := data.Rows[0].Elements[0].Distance.Value / 1000 // In kilometers
    duration := data.Rows[0].Elements[0].Duration.Value / 60  // In minutes

    // Calculate fare
    surge, err := CalculateSurge()
    if err != nil {
        return 0, 0, 0, fmt.Errorf("failed to calculate surge: %v", err)
    }

    fare := CalculateFare(distance, vehicleType, surge)

    return distance, duration, fare, nil
}


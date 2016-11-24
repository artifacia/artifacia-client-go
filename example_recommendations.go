package main

import (
        "github.com/artifacia/go-artifacia/src/client"
)

func main() {
        apiKey := "your_api_key="
        client := client.NewClient(apiKey)
        prodId := "88239"
        num := 20
        filter := `{"color": 1}`
        client.Recommendations.GetVisualRecommendations(prodId, num, filter)
}
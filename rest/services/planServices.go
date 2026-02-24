package services

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/genai"
)

type PlanService struct {
	client *genai.Client
}

func NewPlanService(client *genai.Client) *PlanService {
	return &PlanService{client: client}
}

// GenerateTourPlan generates a travel plan using Gemini and assigns working Picsum placeholder images
func (p *PlanService) GenerateTourPlan(ctx context.Context,
	division string, district string, budget string, locationType string, days int) (map[string]interface{}, error) {

	prompt := fmt.Sprintf(`
You are an expert Bangladesh travel recommendation AI.

User Preferences:
- Division: %s
- District: %s
- Trip Duration: %d days
- Budget Type: %s (economical | normal | luxury)
- Preferred Location Type: %s (chill | nature | urban | mountains)

Tasks:
1. Recommend 4-7 tourist locations in the district.
2. For each location provide:
   - Name
   - Description
   - Activities
   - Rating
   - Accessories
   - Nearby hotels (1-3) with ratings
   - Nearby restaurants (1-3) with ratings
   - Image URL
3. Generate a DAILY ITINERARY:
   - Divide activities among days logically
   - Include weather type for each day (Sunny, Rainy, Cloudy, etc.)
   - Include activities that suit that day's weather
   - Format: "Day 1": {"weather": "Sunny", "activities": ["activity1", "activity2"]}

Consider:
- Weather for next %d days affects activities
- Budget affects hotels/restaurants
- Location type affects activities
- Accessibility and safety

STRICT JSON OUTPUT ONLY:

{
  "division": "",
  "district": "",
  "days": 0,
  "budget": "",
  "location_type": "",
  "weather_considered": true,
  "locations": [
    {
      "name": "",
      "description": "",
      "activities": [],
      "rating": 0,
      "accessories": [],
      "hotels_nearby": [{"name": "", "rating": 0}],
      "restaurants_nearby": [{"name": "", "rating": 0}],
      "image_url": ""
    }
  ],
  "daily_itinerary": {
    "Day 1": {"weather": "", "activities": []},
    "Day 2": {"weather": "", "activities": []}
  }
}
`, division, district, days, budget, locationType, days)

	// Call Gemini API
	result, err := p.client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			ResponseMIMEType: "application/json",
		},
	)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	jsonText := result.Text()
	var output map[string]interface{}
	if err := json.Unmarshal([]byte(jsonText), &output); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v\nraw: %s", err, jsonText)
	}

	// Assign working placeholder images using Lorem Picsum
	if locations, ok := output["locations"].([]interface{}); ok {
		for i, loc := range locations {
			if locMap, ok := loc.(map[string]interface{}); ok {
				locMap["image_url"] = fmt.Sprintf("https://picsum.photos/1600/900?%s&random=%d", locationType, i)
			}
		}
	}

	return output, nil
}

package stars

import (
	"encoding/json"
)

type Star struct {
	Repository     string `json:"full_name"`
	Count          int    `json:"stargazers_count"`
	Link           string `json:"html_url"`
	StarGazersLink string `json:"stargazers_url"`
}

type StarsResponse struct {
	TotalCount int    `json:"total_count"`
	InComplete bool   `json:"incomplete_results"`
	Stars      []Star `json:"items"`
}

func starsFromJson(jsonContent []byte) []Star {
	var result StarsResponse

	json.Unmarshal(jsonContent, &result)

	return result.Stars
}

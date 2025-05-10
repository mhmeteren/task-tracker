package dto

type HealthCheckResponse struct {
	Status  string                 `json:"status"`
	Entries map[string]HealthEntry `json:"entries"`
}

type HealthEntry struct {
	Status string   `json:"status"`
	Tags   []string `json:"tags"`
}

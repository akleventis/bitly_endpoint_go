package models

// API Response models
type User struct {
	Login            string `json:"login"`
	Name             string `json:"name"`
	DefaultGroupGUID string `json:"default_group_guid"`
}

type Bitlinks struct {
	Links []struct {
		Link    string `json:"link"`
		ID      string `json:"id"`
		LongURL string `json:"long_url"`
	} `json:"links"`
}

type CountryMetrics struct {
	Metrics []*Metric `json:"metrics"`
}

type Metric struct {
	Clicks int    `json:"clicks"`
	Value  string `json:"value"`
}

// Client Response models
type AvgClickResponse struct {
	Login         string              `json:"login"`
	GroupGUID     string              `json:"group_guid"`
	TotalBitlinks int                 `json:"num_bitlinks"`
	Metrics       map[string]*Average `json:"averages"`
}
type Average struct {
	AverageClicks float64 `json:"average_clicks"`
}

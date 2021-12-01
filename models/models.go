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
	Unit          string `json:"unit"`
	Units         int    `json:"units"`
	Facet         string `json:"facet"`
	UnitReference string `json:"unit_reference"`
	Metrics       []struct {
		Clicks int    `json:"clicks"`
		Value  string `json:"value"`
	} `json:"metrics"`
}

// Client Response models
type avgClickResponse struct {
	Login         string `json:"login"`
	GroupGUID     string `json:"group_guid"`
	TotalBitlinks int    `json:"num_bitlinks"`
	Metrics       []struct {
		TotalClicks   int     `json:"total_clicks"`
		AverageClicks float64 `json:"average_clicks"`
	}
}

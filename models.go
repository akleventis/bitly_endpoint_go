package main

type User struct {
	Login     string `json:"login"`
	Name      string `json:"name"`
	IsActive  bool `json:"is_active"`
	Created   string `json:"created"`
	Modified  string `json:"modified"`
	IsSsoUser bool `json:"is_sso_user"`
	Emails    []struct {
		Email      string `json:"email"`
		IsPrimary  bool `json:"is_primary"`
		IsVerified bool `json:"is_verified"`
	} `json:"emails"`
	Is2FaEnabled     bool `json:"is_2fa_enabled"`
	DefaultGroupGUID string `json:"default_group_guid"`
}

type Bitlinks struct {
	Links []struct {
		References struct {
			Group string `json:"group"`
		} `json:"references"`
		Link           string   `json:"link"`
		ID             string   `json:"id"`
		LongURL        string   `json:"long_url"`
		Title          string   `json:"title"`
		Archived       bool   `json:"archived"`
		CreatedAt      string   `json:"created_at"`
		CreatedBy      string   `json:"created_by"`
		ClientID       string   `json:"client_id"`
		CustomBitlinks []string `json:"custom_bitlinks"`
		Tags           []string `json:"tags"`
		LaunchpadIds   []string `json:"launchpad_ids"`
		Deeplinks      []struct {
			GUID        string `json:"guid"`
			Bitlink     string `json:"bitlink"`
			AppURIPath  string `json:"app_uri_path"`
			InstallURL  string `json:"install_url"`
			AppGUID     string `json:"app_guid"`
			Os          string `json:"os"`
			InstallType string `json:"install_type"`
			Created     string `json:"created"`
			Modified    string `json:"modified"`
			BrandGUID   string `json:"brand_guid"`
		} `json:"deeplinks"`
	} `json:"links"`
	Pagination struct {
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Size  int `json:"size"`
		Page  int `json:"page"`
		Total int `json:"total"`
	} `json:"pagination"`
}

	
type CountryMetrics struct {
	Unit          string `json:"unit"`
	Units         int `json:"units"`
	Facet         string `json:"facet"`
	UnitReference string `json:"unit_reference"`
	Metrics       []struct {
		Clicks int `json:"clicks"`
		Value  string `json:"value"`
	} `json:"metrics"`
}
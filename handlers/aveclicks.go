package handlers

import (
	"bitly_server_go/client"
	"encoding/json"
	"errors"
	"math"
	"net/http"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) GetClicks(w http.ResponseWriter, r *http.Request) {
	authToken, err := getAuthToken(r)
	if err != nil {
		Error(w, 403, err)
	}
	client := client.New(authToken)

	// if group guid is provided, use it, if not, grab defualt from user request
	var groupGuid string
	switch guid := r.URL.Query().Get("groupGuid"); guid {
	case "":
		user := client.GetUser()
		groupGuid = user.DefaultGroupGUID
	default:
		groupGuid = guid
	}

	allLinks := client.GetLinks(groupGuid)

	// create array of bitlink id's
	bitlinkArr := []string{}
	for _, obj := range allLinks.Links {
		bitlinkArr = append(bitlinkArr, obj.ID)
	}

	// eventually return this map
	countries := make(map[string]float64)
	unit, units := "month", "30"

	for _, link := range bitlinkArr {
		data := client.GetClicksByCountry(link, unit, units)

		// loop through clicks (struct[] of k:countries v:clicks), if no click data, len is 0, continue to next link
		for i := 0; i < len(data.Metrics); i++ {
			clicks, country := data.Metrics[i].Clicks, data.Metrics[i].Value

			// if country in map, add clicks to value, else initialize country with current link clicks divided by 30 days
			if _, ok := countries[country]; ok {
				countries[country] += (float64(clicks) / 30)
			} else {
				countries[country] = (float64(clicks) / 30)
			}
		}
	}
	// Round off numbers
	for k, v := range countries {
		countries[k] = math.Round(v*1000) / 1000
	}

	// make response map => nest countries[map] inside
	respoonseMap := make(map[string]map[string]float64)
	respoonseMap["Average daily clicks per country over the past month"] = countries

	json.NewEncoder(w).Encode(respoonseMap)

}

func getAuthToken(r *http.Request) (string, error) {
	token := r.Header["Authorization"][0]
	if token == "" {
		return "", errors.New("Missing/Invalid auth header")
	}
	return token, nil
}

func Error(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}

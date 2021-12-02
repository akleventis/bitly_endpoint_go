package handlers

import (
	"bitly_server_go/client"
	"bitly_server_go/models"
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

	// if group guid is provided, use it, if not, grab default from user request
	groupGuid, user := getGroupGuid(r, client)

	allLinks := client.GetLinks(groupGuid)

	// create array of bitlink id's
	bitlinkArr := []string{}
	for _, obj := range allLinks.Links {
		bitlinkArr = append(bitlinkArr, obj.ID)
	}

	countryData := buildCountryAve(client, bitlinkArr)

	averages := make(map[string]*models.Average)
	for k, v := range countryData {
		country := &models.Average{
			AverageClicks: v,
		}
		averages[k] = country
	}

	response := &models.AvgClickResponse{
		Login:         user.Login,
		GroupGUID:     groupGuid,
		TotalBitlinks: len(bitlinkArr),
		Metrics:       averages,
	}
	json.NewEncoder(w).Encode(response)
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

func getGroupGuid(r *http.Request, c *client.Auth) (string, models.User) {
	var groupGuid string
	user := c.GetUser()
	switch guid := r.URL.Query().Get("groupGuid"); guid {
	case "":
		groupGuid = user.DefaultGroupGUID
	default:
		groupGuid = guid
	}
	return groupGuid, user
}

func buildCountryAve(client *client.Auth, bitlinkArr []string) map[string]float64 {
	countries := make(map[string]float64)
	unit, units := "month", "30"

	for _, link := range bitlinkArr {
		data := client.GetClicksByCountry(link, unit, units)
		// loop through clicks (struct[] of k:countries v:clicks), if no click data, len is 0, continue to next link
		for i := 0; i < len(data.Metrics); i++ {
			clicks, country := data.Metrics[i].Clicks, data.Metrics[i].Value
			// if country in map, add clicks to value, else initialize country with current clicks divided by 30 days
			if _, ok := countries[country]; ok {
				countries[country] += (float64(clicks) / 30)
			} else {
				countries[country] = (float64(clicks) / 30)
			}
		}
	}
	for k, v := range countries {
		countries[k] = math.Round(v*1000) / 1000
	}
	return countries
}

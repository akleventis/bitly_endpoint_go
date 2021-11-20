package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
)

var authToken string;

func getClicks(w http.ResponseWriter, r *http.Request) {
	authToken = r.Header["Authorization"][0]

	user := getUser()

	groupGuid := user.DefaultGroupGUID

	allLinks := getLinks(groupGuid)

	bitlinkArr := []string{}
	
	for _, obj := range allLinks.Links{
		bitlinkArr = append(bitlinkArr, obj.ID)
	}
	countries := make(map[string]float64)
	unit, units := "month", "30"

	for _, link := range bitlinkArr {
		data := getClicksByCountry(link, unit, units)

		var clicks int
		var country string

		// loop through clicks (struct[] of k:countries v:clicks)
		// if no click data, len is 0, continue to next link
		for i:=0; i<len(data.Metrics);i++ {
			clicks = data.Metrics[i].Clicks
			country = data.Metrics[i].Value
			// if country in map, add clicks to value, else initialize country with current link clicks
			if _, ok := countries[country]; ok {
				countries[country] += float64(clicks)
			} else {
				countries[country] = float64(clicks)
			}
		}
	}
	// divide by 30, round 3 places
	for k := range countries{
		x := countries[k]/30
		countries[k] = math.Round(x*1000)/1000
	}
	
	// make response map => nest countries[map] inside
	respoonseMap := make(map[string]map[string]float64)
	respoonseMap["Average daily clicks per country over the past month"] = countries

	json.NewEncoder(w).Encode(respoonseMap)

}

func handleRequests() {
	http.HandleFunc("/clicks", getClicks)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func main() {
	handleRequests()
}


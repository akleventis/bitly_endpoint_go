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
	
	var groupGuid string;
	
	// if group guid is provided, use it, if not, grab defualt from user request
	switch guid := r.URL.Query().Get("groupGuid"); guid {
	case "":
		user := getUser()
		groupGuid = user.DefaultGroupGUID
	default:
		groupGuid = guid
	}

	allLinks := getLinks(groupGuid)

	// create array of bitlink id's
	bitlinkArr := []string{}
	for _, obj := range allLinks.Links{
		bitlinkArr = append(bitlinkArr, obj.ID)
	}

	// eventually return this map
	countries := make(map[string]float64)
	unit, units := "month", "30"

	for _, link := range bitlinkArr {
		data := getClicksByCountry(link, unit, units)

		var clicks int
		var country string

		// loop through clicks (struct[] of k:countries v:clicks)
		// if no click data, len is 0, continue to next link
		for i:=0; i<len(data.Metrics);i++ {
			clicks, country = data.Metrics[i].Clicks, data.Metrics[i].Value
			
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


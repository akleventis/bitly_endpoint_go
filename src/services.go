package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getUser() User {
	url := "https://api-ssl.bitly.com/v4/user"
	body, err := retrieveData(url)
	if err != nil {
		log.Fatalln(err)
	}
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatalln(err)
	}
	return user
}

func getLinks(group_guid string) Bitlinks {
	url := fmt.Sprintf("https://api-ssl.bitly.com/v4/groups/%s/bitlinks", group_guid)
	body, err := retrieveData(url)
	if err != nil {
		log.Fatalln(err)
	}
	var bitlinks Bitlinks
	if err := json.Unmarshal(body, &bitlinks); err != nil{
		log.Fatalln(err)
	}
	return bitlinks
}

func getClicksByCountry(bitlink string, unit string, units string) CountryMetrics {
	url := fmt.Sprintf("https://api-ssl.bitly.com/v4/bitlinks/%s/countries?day=%s&units=%s", bitlink, unit, units)
	body, err := retrieveData(url)
	if err != nil {
		log.Fatalln(err)
	}
	var clicksByCountry CountryMetrics
	if err := json.Unmarshal(body, &clicksByCountry); err != nil {
		log.Fatalln(err)
	}
	return clicksByCountry
}

func retrieveData(url string) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
		"Authorization": []string{authToken},
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}




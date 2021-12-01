package client

import (
	models "bitly_server_go/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const URL = "https://api-ssl.bitly.com/v4"

type Auth struct {
	AuthToken string
}

func New(authToken string) *Auth {
	return &Auth{
		AuthToken: authToken,
	}
}

func (a *Auth) GetUser() models.User {
	url := "https://api-ssl.bitly.com/v4/user"
	body, err := a.retrieveData(url)
	if err != nil {
		log.Fatalln(err)
	}
	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatalln(err)
	}
	return user
}

func (a *Auth) GetLinks(group_guid string) models.Bitlinks {
	url := fmt.Sprintf("https://api-ssl.bitly.com/v4/groups/%s/bitlinks", group_guid)
	body, err := a.retrieveData(url)
	if err != nil {
		log.Fatalln(err)
	}
	var bitlinks models.Bitlinks
	if err := json.Unmarshal(body, &bitlinks); err != nil {
		log.Fatalln(err)
	}
	return bitlinks
}

func (a *Auth) GetClicksByCountry(bitlink string, unit string, units string) models.CountryMetrics {
	url := fmt.Sprintf("https://api-ssl.bitly.com/v4/bitlinks/%s/countries?day=%s&units=%s", bitlink, unit, units)
	body, err := a.retrieveData(url)
	if err != nil {
		log.Fatalln(err)
	}
	var clicksByCountry models.CountryMetrics
	if err := json.Unmarshal(body, &clicksByCountry); err != nil {
		log.Fatalln(err)
	}
	return clicksByCountry
}

// Helper function: Returns []byte response from provided URL
func (a *Auth) retrieveData(url string) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{a.AuthToken},
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

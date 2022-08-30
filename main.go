package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/kr/pretty"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Fedex struct {
		ID             string
		Secret         string
		URL            string
		TrackingNumber string
	}
}

func main() {
	// config
	data, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err = toml.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}

	// oauth token
	v := url.Values{}
	v.Set("grant_type", "client_credentials")
	v.Set("client_id", config.Fedex.ID)
	v.Set("client_secret", config.Fedex.Secret)

	resp, err := http.PostForm(config.Fedex.URL+"/oauth/token", v)
	if err != nil {
		log.Fatal(err)
	}

	var oauth OAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&oauth); err != nil {
		log.Fatal(err)
	}

	fmt.Println(oauth)

	// tracking
	payload := TrackingRequest{
		IncludeDetailedScans: true,
		TrackingInfo: []TrackingInfoRequest{
			{
				TrackingNumberInfo: TrackingNumberInfoRequest{
					TrackingNumber: config.Fedex.TrackingNumber,
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, config.Fedex.URL+"/track/v1/trackingnumbers", &buf)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-locale", "en_US")
	req.Header.Set("Authorization", "Bearer "+oauth.AccessToken)

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var tracking TrackResponse
	if err := json.NewDecoder(resp.Body).Decode(&tracking); err != nil {
		log.Fatal(err)
	}

	pretty.Println(tracking)
}

package fedex

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type Client struct {
	Config Config
	OAuth  OAuthResponse
}

type Config struct {
	ID     string
	Secret string
	URL    string
}

func (c *Client) RefreshToken() error {
	v := url.Values{}
	v.Set("grant_type", "client_credentials")
	v.Set("client_id", c.Config.ID)
	v.Set("client_secret", c.Config.Secret)

	resp, err := http.PostForm(c.Config.URL+"/oauth/token", v)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(&c.OAuth); err != nil {
		return err
	}

	return nil
}

func (c *Client) Track(number string) (tracking TrackResponse, err error) {
	if c.OAuth == (OAuthResponse{}) {
		if err = c.RefreshToken(); err != nil {
			return
		}
	}

	payload := TrackingRequest{
		IncludeDetailedScans: true,
		TrackingInfo: []TrackingInfoRequest{
			{
				TrackingNumberInfo: TrackingNumberInfoRequest{
					TrackingNumber: number,
				},
			},
		},
	}

	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(payload); err != nil {
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, c.Config.URL+"/track/v1/trackingnumbers", &buf)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-locale", "en_US")
	req.Header.Set("Authorization", "Bearer "+c.OAuth.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if err = json.NewDecoder(resp.Body).Decode(&tracking); err != nil {
		return
	}

	return
}

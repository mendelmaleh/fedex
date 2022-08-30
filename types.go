package fedex

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type TrackingRequest struct {
	IncludeDetailedScans bool                  `json:"includeDetailedScans"`
	TrackingInfo         []TrackingInfoRequest `json:"trackingInfo"`
}

type TrackingInfoRequest struct {
	TrackingNumberInfo TrackingNumberInfoRequest `json:"trackingNumberInfo"`
}

type TrackingNumberInfoRequest struct {
	TrackingNumber string `json:"trackingNumber"`
}

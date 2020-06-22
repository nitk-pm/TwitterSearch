package model

type OAuth struct {
	ConsumerKey    string
	ConsumerSecret string
	OAuthToken     string
	OAuthSecret    string
}

type TokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}
type SearchResponse struct {
	Statuses []struct {
		Text string `json:"text"`
	} `json:"statuses"`
}

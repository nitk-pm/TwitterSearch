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

// type SearchResponse struct {
// 	Statuses []struct {
// 		CreatedAt string `json:"created_at"`
// 		Id        string `json:"id_str"`
// 		Text      string `json:"text"`
// 	} `json:"statuses"`
// }

type SearchResponse struct {
	Statuses []map[string]interface{} `json:"statuses"`
}

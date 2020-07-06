package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"repEATer/model"
)

func SearchTweets(accessToken string, queryParam string, url string) (*model.SearchResponse, error) {
	req, err := http.NewRequest("GET", url+"?q="+queryParam, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if err := HTTPResponseCheck(res.StatusCode); err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", res) //Search APIのレスポンス確認
	defer res.Body.Close()
	searchResponse := &model.SearchResponse{}
	body, _ := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(body, searchResponse); err != nil {
		return nil, err
	}
	return searchResponse, nil
}

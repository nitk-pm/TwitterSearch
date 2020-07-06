package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"repEATer/model"
)

func GetToken(oauth model.OAuth, url string) (string, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBufferString("grant_type=client_credentials"))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Basic "+MakeCredential(oauth))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	res, err := http.DefaultClient.Do(req)
	// fmt.Printf("%+v\n", res) //トークンリクエストのレスポンス確認
	if err != nil {
		return "", err
	}
	err = HTTPResponseCheck(res.StatusCode)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	tokenResponse := &model.TokenResponse{}
	body, _ := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(body, tokenResponse); err != nil {
		return "", err
	}
	return tokenResponse.AccessToken, nil
}

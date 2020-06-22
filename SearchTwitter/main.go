package main

import (
	"./handler"
	"./model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const bearerTokenURL = "https://api.twitter.com/oauth2/token"
const resourceURL = "https://api.twitter.com/1.1/search/tweets.json"

//TODO 頼むからリファクタリングしてくれ．
func main() {
	var oauth model.OAuth
	oauth.ConsumerKey = os.Getenv("REPEATER_CONSUMER_KEY")
	oauth.ConsumerSecret = os.Getenv("REPEATER_CONSUMER_SECRET")

	if oauth.ConsumerKey == "" || oauth.ConsumerSecret == "" {
		fmt.Println("環境変数が設定されていません．REPEATER_CONSUMER_(KEY|SECRET)にAPI KeyとSecretの値を設定してください")
		return
	}

	//GetTokenとかに切り分けよう
	req, _ := http.NewRequest("POST", bearerTokenURL, bytes.NewBufferString("grant_type=client_credentials"))
	req.Header.Add("Authorization", "Basic "+handler.MakeCredential(oauth))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	res, _ := http.DefaultClient.Do(req)
	// fmt.Printf("%+v\n", res) //トークンリクエストのレスポンス確認
	if res.StatusCode != 200 {
		fmt.Printf("正常に処理されませんでした エラーコード:%d\n", res.StatusCode)
		switch res.StatusCode {
		case 403:
			fmt.Println("多分API KeyかSecretが間違っています")
		case 404:
			fmt.Println("URLのTypoやAPI側のリソースURLの変更を確認してください")
		case 420, 429:
			fmt.Printf("クエリ送りすぎです ")
		}
		fmt.Println("https://developer.twitter.com/ja/docs/basics/response-codes で詳細を確認してください")
		return
	}
	defer res.Body.Close()
	tokenResponse := &model.TokenResponse{}
	body, _ := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(body, tokenResponse); err != nil {
		fmt.Println("JSONパースエラー:", err)
	}
	accessToken := tokenResponse.AccessToken

	queryParam := url.QueryEscape("from:@nosykcam")
	req, _ = http.NewRequest("GET", resourceURL+"?q="+queryParam, nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)

	//TODO res2とかいうバカみたいな変数名をどうにかしろ
	res2, _ := http.DefaultClient.Do(req)
	// fmt.Printf("%+v\n", res2) //Search APIのレスポンス確認
	if res2.StatusCode != 200 {
		fmt.Printf("正常に処理されませんでした エラーコード:%d\n", res.StatusCode)
		switch res.StatusCode {
		case 403:
			fmt.Println("多分トークンが間違っています")
		case 404:
			fmt.Println("URLのTypoやAPI側のリソースURLの変更を確認してください")
		case 406:
			fmt.Println("検索クエリの中身に不正な値が含まれています")
		case 420, 429:
			fmt.Printf("クエリ送りすぎです ")
		}
		fmt.Println("https://developer.twitter.com/ja/docs/basics/response-codes で詳細を確認してください")
		return
	}
	defer res2.Body.Close()
	searchResponse := &model.SearchResponse{}
	body, _ = ioutil.ReadAll(res2.Body)
	if err := json.Unmarshal(body, searchResponse); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(searchResponse.Statuses[0].Text)
}

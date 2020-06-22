package main

import (
	"./handler"
	"./model"
	"fmt"
	"log"
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
		log.Fatal("環境変数が設定されていません．REPEATER_CONSUMER_(KEY|SECRET)にAPI KeyとSecretの値を設定してください")
	}

	accessToken, err := handler.GetToken(oauth, bearerTokenURL)
	if err != nil {
		log.Fatal(err)
	}
	queryParam := url.QueryEscape("from:@nosykcam")
	searchResponse, err := handler.SearchTweets(accessToken, queryParam, resourceURL)
	if err != nil {
		log.Fatal(err)
	}

	//とりあえず拾ったツイート全列挙してみる
	for _, tweet := range searchResponse.Statuses {
		fmt.Printf("%s\n", tweet.Text)
	}
}

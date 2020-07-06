package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"repEATer/db"
	"repEATer/handler"
	"repEATer/model"
)

const bearerTokenURL = "https://api.twitter.com/oauth2/token"
const resourceURL = "https://api.twitter.com/1.1/search/tweets.json"

func main() {
	es, err := db.GetElasticsearchClient()
	if err != nil {
		log.Fatal(err)
	}
	es.Info() //まだデータは投げないのでこれは無駄撃ちです．

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
		txt, err := json.MarshalIndent(tweet, "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(txt), "\n")
	}
	// for _, tweet := range searchResponse.Statuses {
	// 	db.AddData(es, tweet)
	// }
}

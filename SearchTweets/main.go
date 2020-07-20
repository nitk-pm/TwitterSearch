package main

import (
	"TwitterSearch/db"
	"TwitterSearch/handler"
	"TwitterSearch/model"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
)

const bearerTokenURL = "https://api.twitter.com/oauth2/token"
const resourceURL = "https://api.twitter.com/1.1/search/tweets.json"

func main() {
	query := flag.String("query", "", "search query")
	count := flag.Int("count", 15, "number of obtained tweets")
	flag.Parse()

	if *query == "" {
		log.Fatal("クエリが指定されていません\n ----------------------USAGE----------------------\n./TweetSearch -query=from:@twitter -count=15\n--------------------------------------------------\n\n")
	}

	es, err := db.GetDBClient()
	if err != nil {
		log.Fatal(err)
	}

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
	queryParam := url.QueryEscape(*query)

	for {
		searchResponse, err := handler.SearchTweets(accessToken, queryParam, *count, resourceURL)
		if err != nil {
			log.Fatal(err)
		}

		err = db.AddData(es, searchResponse)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Done")
		time.Sleep(5 * time.Hour) //一日50件ないぐらいなので、５時5ごとに百100件回せば間違いなく漏れはない。
	}
}

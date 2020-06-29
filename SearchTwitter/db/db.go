// package db
package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}
	text := "{\"title\":\"test3\"}"
	res, err := AddData(es, text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
func AddData(es *elasticsearch.Client, text string) (string, error) { //副作用あります，しょうがないよね．
	req := esapi.IndexRequest{
		Index: "test",
		Body:  strings.NewReader(text),
	}
	res, err := req.Do(context.Background(), es)
	return res.String(), err
}

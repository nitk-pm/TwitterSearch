// package db
package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	// "github.com/elastic/go-elasticsearch/v8/esapi"
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
	res, err := es.Index(
		"test", //Index(データベース)の名前
		strings.NewReader("{\"title\":\"test2\"}"), //Document(レコード)はReader型でないとだめ(多分クソデカいデータ入れることもあるため)
	) //"_type"は_docで固定，"_id"はデータ記録時にランダム生成
	return res.String(), err
}

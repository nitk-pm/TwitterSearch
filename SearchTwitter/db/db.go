package db

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"strings"
)

func GetElasticsearchClient() (*elasticsearch.Client, error) {
	es, err := elasticsearch.NewDefaultClient()
	return es, err
}
func AddData(es *elasticsearch.Client, text string) (string, error) { //副作用あります，しょうがないよね．
	req := esapi.IndexRequest{
		Index: "test",
		Body:  strings.NewReader(text),
	}
	res, err := req.Do(context.Background(), es)
	return res.String(), err
}

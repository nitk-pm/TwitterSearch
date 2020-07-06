package db

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"repEATer/model"
	"strings"
)

func GetDBClient() (*elasticsearch.Client, error) {
	es, err := elasticsearch.NewDefaultClient()
	return es, err
}

//TODO ESの投げたエラーをハンドリングしていないので，動いたのにデータが取れてないみたいなことになりかねない．
//これはバルク処理にできるはず?
func AddData(es *elasticsearch.Client, searchResponse *model.SearchResponse) error { //stringを返すべきか謎
	deleteDuplicateTweetInfo(es, searchResponse)
	for _, tweetInfo := range (*searchResponse).Statuses {
		tweetInfoJSONText, err := json.Marshal(tweetInfo)
		if err != nil {
			return err
		}
		req := esapi.IndexRequest{
			Index: "tweet",
			Body:  strings.NewReader(string(tweetInfoJSONText)),
		}
		_, err = req.Do(context.Background(), es)
		// res, err := req.Do(context.Background(), es) //resを受け取ってハンドルする
		if err != nil {
			return err
		}
	}
	return nil
}
func deleteDuplicateTweetInfo(es *elasticsearch.Client, searchResponse *model.SearchResponse) error {
	// 一つひとつのツイートに対してID被りがあるかどうかチェックするよりも，突っ込むツイートのID全部のOR検索でマッチしたやつ全部消したほうがリクエストの数が減って効率いいと思うので，そんな実装します．多分HTTP通信がボトルネックなので．
	idList := make([]string, 0)
	for _, tweetInfo := range (*searchResponse).Statuses {
		idList = append(idList, tweetInfo["id_str"].(string))
	}
	idListText := strings.Join(idList[:], ",")
	_ = idListText
	query := "{\"query\": {\"terms\": {\"id_str\":[" + idListText + "]}}}"
	req := esapi.DeleteByQueryRequest{
		Index: []string{"tweet"},
		Body:  strings.NewReader(query),
	}
	_, err := req.Do(context.Background(), es)
	return err
}

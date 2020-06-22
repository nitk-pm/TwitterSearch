package handler

import (
	"fmt"
)

func HTTPResponseCheck(statusCode int) error {
	// var err error //TODO 標準出力脱却する
	if statusCode != 200 {
		fmt.Printf("正常に処理されませんでした エラーコード:%d\n", statusCode)
		switch statusCode {
		case 403:
			fmt.Println("多分API Key，Secret，トークンが間違っています")
		case 404:
			fmt.Println("URLのTypoやAPI側のリソースURLの変更を確認してください")
		case 420, 429:
			fmt.Printf("クエリ送りすぎです ")
		}
		fmt.Println("https://developer.twitter.com/ja/docs/basics/response-codes で詳細を確認してください")
		return fmt.Errorf("ErrorCode:%d", statusCode)
	}
	return nil
}

package handler

import (
	"fmt"
)

func HTTPResponseCheck(statusCode int) error {
	if statusCode != 200 {
		msg := fmt.Sprintf("正常に処理されませんでした エラーコード:%d\n", statusCode)
		switch statusCode {
		case 403:
			msg += fmt.Sprintf("多分API Key，Secret，トークンが間違っています\n")
		case 404:
			msg += fmt.Sprintf("URLのTypoやAPI側のリソースURLの変更を確認してください\n")
		case 420, 429:
			msg += fmt.Sprintf("クエリ送りすぎです\n")
		}
		msg += ("https://developer.twitter.com/ja/docs/basics/response-codes で詳細を確認してください")
		return fmt.Errorf(msg)
	}
	return nil
}

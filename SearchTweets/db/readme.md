# Elasticsearchについて
検索が早くて，JSONデータを入れるのが簡単なデータベースです．

ローカルで使う場合，Elasticsearch起動までは各自で調べてください．

そのうち動かなくなるかもですが，その時はここだけ改修してください．ツイート検索はGoの標準パッケージだけで作っているので大丈夫．
## データの操作方法
ラッパーを使わない場合，所定のURL(なにもいじってない場合<http://localhost:9200>/hoge)に対して(GET|POST)リクエストを送ることでCRUD操作ができます(Index,Get,Update,Delete APIが対応する)．[リファレンスはここ](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs.html)

ラッパーを使う場合はそのラッパーのドキュメントを読みましょう．

(Golangは公式のクライアントがあります．[リポジトリ](https://github.com/elastic/go-elasticsearch)・[ドキュメント](https://pkg.go.dev/github.com/elastic/go-elasticsearch/esapi))

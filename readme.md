# Twitter検索ツール
## 開発環境
* Golang v1.14.4
* Elasticsearch v7.8.0
* Kibana v7.8.0

完璧に同じ環境じゃなくても全然動くと思いますが，困ったらバージョンを確認してください．

## 使い方
実行前にdocker-composeを入手してください．
```
$ sudo docker-compose up -d
$ cd ./SearchTweets && go build -o ../Search && cd ..
$ ./Search -query=from:@Twitter -count=20
```

これで@Twitterのツイートを20件取得できます．countは最大100までです．

取得したデータはlocalhost:5601にブラウザでアクセスして，メニューの中にあるDiscoverのページで見ることが出来ます(Kibanaの起動は遅いので注意)．
## (DB以外を)改造する際の前提知識
* Twitter APIのKeyとSecretを取得できる
* ベアラートークンを使ってSearch APIを叩ける
* Twitter APIドキュメント(英語)を読める
* docker-compose.ymlに何が書いているかわかる

DBについては./SearchTweets/db/readme.mdを読んでください．
## 開発フロー
1. nitk-pm/Twittersearch(upstreamという)をclone
1. cloneしたリポジトリにコミット
1. GitHub上からupstreamへPull Requestを出す
1. 修正が必要なら修正し，マージする

プルリクの#3に単純な例を置きました．

## バグの修正
1. nitk-pm/TwitterSearchにIssueを立てる
1. 調査のために適当にブランチを切って実験する
1. 原因が分かったら調査用ブランチを消し，開発用ブランチからfeature/#N\_bug\_descriptionというブランチを切る
1. 直ったら開発用ブランチへマージし，upstreamへプルリクを出す

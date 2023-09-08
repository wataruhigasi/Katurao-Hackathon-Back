# ローカルでサーバーを立ち上げる方法

```
$ docker-compose up --build
```

## リクエストとレスポンス
```
記事作成
- リクエスト
$ curl -X POST -H "Content-Type: application/json" -d '{
  "title": "New Article",
  "body": "This is the body of the new article.",
  "position": {
    "x": 35,
    "y": 45
  }
}' http://localhost:8080/article

```

```
記事取得
- リクエスト
$ curl http://localhost:8080/articles | jq
```
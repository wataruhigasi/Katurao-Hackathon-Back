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

- レスポンス
[
  {
    "id": 1,
    "title": "Article 1",
    "created_at": "2023-09-08T09:50:01Z",
    "body": "This is the body of Article 1.",
    "position": {
      "X": 10,
      "Y": 20
    }
  },
  {
    "id": 2,
    "title": "Article 2",
    "created_at": "2023-09-08T09:50:01Z",
    "body": "This is the body of Article 2.",
    "position": {
      "X": 15,
      "Y": 25
    }
  },
]

```
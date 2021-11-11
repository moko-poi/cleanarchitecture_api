### go による RESTAPIの実装（cleanarchitecture）

### 内容
- /api/user/register(userの登録　POST)
- /api/user/get(userの全取得　GET)
- /api/user/get/:id(useridからの取得 GET)

・目的として、クリーンアーキテクチャの理解、実装
ユーザーリソースを提供する単純なアプリケーションです。
・POST /users ユーザー登録
・GET /users ユーザー一覧取得
・GET /user/:id ユーザー情報の取得
・Userは、id firstName lastNameを持つ。

### 使用技術
- go
- cleanarchitecture
- gorm
- MySQL

### 使用方法
MySQLサーバーの立ち上げ
```
docker-compose up -d
```

api
```
go run main.go
```

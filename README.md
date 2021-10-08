### go による RESTAPIの実装（cleanarchitecture）

### 内容
- /api/user/register(userの登録　POST)
- /api/user/get(userの全取得　GET)
- /api/user/get/:id(useridからの取得 GET)

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

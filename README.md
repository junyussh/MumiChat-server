# MumiChat-Server
---

這個專案是學校暑假實訓所做的項目，使用 Golang 和 gin 進行後端開發，通訊協議採用 Websocket 連線，未來可能會增加其他 API 接口，

## 自行編譯(Build)

此步驟是為欲自行編譯者提供，如果是使用釋出的二進制執行檔者可跳過此步。

請先將 Golang 執行環境安裝好，安裝方法請參考[官方文檔](https://golang.org/doc/install)，然後將 repo clone 到 `$GOPATH`。

```bash
git clone https://github.com/junyussh/MumiChat-server.git
```

`$GOPATH` 路徑可在終端機輸入 `go env` 獲取 Golang 環境變數。

然後用 `go get` 安裝依賴套件。

```bash
go get
```

## 使用(Usage)

### 設定

`conf/app.ini` 目前可以設定 gin 的輸出模式，資料庫儲存位置、名稱、還有伺服器監聽的端口號，主要就這幾個選項。

```ini
# possible values : release, debug
app_mode = debug

[paths]
# Path to where grafana can store temp files, sessions, and the sqlite3 db (if that is used)
data = ./data

[server]
# The http port  to use
HTTP_PORT = 8080

[database]
TYPE = sqlite3
NAME = chat
PATH = ./
TABLE_PREFIX = chat
```

### 執行

執行 `main.go`。

```bash 
go run main.go
```

如果使用二進制執行檔者就直接執行。

```bash
./main
```

## 請求標準

當連接上伺服器，伺服器會回傳一串 key，為伺服器的公鑰，未來要做非對稱加密連線用。

```json
{
  "key": "24AB50C2"
}
```

每個請求都有個 `type` 欄，目前有兩種值：`action` 和 `message`，註冊登入屬於 `action`，而訊息則是 `message`。

`action`: 合法值目前只有 `register` 和 `login`，分別代表註冊和登入請求。

```json
{
    "type": "action",
    "action": "login"
}
```

### 註冊

註冊帳號必填以下欄位資訊：

- `email`: 使用者 Email
- `password`: 密碼
- `firstName`: 名字
- `lastName`: 姓氏

```json
{
  "type": "action",
  "action": "register",
  "data": {
    "email": "abc@example.com",
    "password": "p@ssw0rd",
    "firstName": "Eric",
    "lastName": "Chen"
  }
}
```

### 登入

登入只須填 Email 和密碼，`action` 欄位的值改為 `login` 即可。

註：目前設計一個帳號僅能有一個 Websocket 連線。

```json
{
  "type": "action",
  "action": "login",
  "data": {
    "email": "abc@example.com",
    "password": "p@ssw0rd"
  }
}
```

## 發送訊息

發送訊息前必須先進行登入。

訊息欄位須包含：

- `type`: 這裡必須填 `message`
- `recipient`: 接收者的 ID
- `content`: 訊息內容

整體請求如下：

```json
{
  "type": "message",
  "data": {
    "recipient": "289559048",
    "content": "hello"
  }
}
```

發送成功後會回傳：

```json
{
  "code": 201,
  "msg": "訊息發送成功",
  "data": null
}
```

然後接收者會收到發送者傳來的訊息，`sender` 為訊息發送者，`created_at` 為訊息發送時間。

```json
{
  "sender": "271693512",
  "recipient": "289559048",
  "content": "hello",
  "created_at": "2019-07-15T09:28:10Z"
}
```

## License

The project is under MIT License now.
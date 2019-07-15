package e

// MsgFlags maps error code to error message
var MsgFlags = map[int]string {
	SUCCESS: "ok",
	MESSAGE_SUCCESS: "訊息發送成功",
	ERROR: "failed",
	ERROR_AUTH: "使用者名稱或密碼錯誤",
	ERROR_AUTH_FAILED: "認證失敗請重試",
	ERROR_ADD_USER_FAILED: "使用者創建失敗",
	ERROR_EXIST_USER: "使用者已存在",
	ERROR_EXIST_USER_FAILED: "查詢使用者失敗",
	ERROR_MULTIPLE_LOGIN: "使用者已在其他地方登入",
	ERROR_MULTIPLE_LOGIN_FAILED: "使用者登入狀態獲取失敗",
	ERROR_UNAUTHORIZED: "請先進行登入",
	ERROR_USER_NOT_ONLINE: "使用者目前不在線上",
	ERROR_CONNECTION_CLOSED: "連線已關閉",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

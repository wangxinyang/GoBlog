package code

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_USERNAME_USED = 1001
)

var codeMap = map[int]string{
	SUCCESS:             "操作成功",
	ERROR:               "操作失敗",
	ERROR_USERNAME_USED: "当該ユーザーが存在しました。",
}

// GetMsg コードによって、メッセージを戻す
func GetMsg(code int) string {
	return codeMap[code]
}

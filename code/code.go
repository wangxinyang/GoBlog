package code

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_USERNAME_USED = 1001

	ERROR_CATEGORYNAME_USED = 2001
)

var codeMap = map[int]string{
	SUCCESS:                 "操作成功",
	ERROR:                   "操作失敗",
	ERROR_USERNAME_USED:     "当該ユーザーが存在しました。",
	ERROR_CATEGORYNAME_USED: "当該カテゴリが存在しました。",
}

// GetMsg コードによって、メッセージを戻す
func GetMsg(code int) string {
	return codeMap[code]
}

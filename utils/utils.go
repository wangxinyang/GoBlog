package utils

import (
	"crypto/md5"
	"fmt"
)

// GetMD5 MD5の文字を取得する
func GetMD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

// InList 判断key是否存在列表
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

// Md5 md5
func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

func IsList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false

}

func FormatTime(time time.Time) string {
	return time.Format("2006-01-02 15:04")
}

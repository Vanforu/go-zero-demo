package utils

import (
	"crypto/sha256"
	"fmt"
)

// 加密
func Encrypt(data string, salt string) string {
	r := sha256.Sum256([]byte(data + "$" + salt))
	r_str := fmt.Sprintf("%x", r)
	return r_str
}

// 对比
func Compare(data string, password string, salt string) bool {
	r := sha256.Sum256([]byte(data + "$" + salt))
	r_str := fmt.Sprintf("%x", r)
	return r_str == password
}

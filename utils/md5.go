package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// md5加密
func Md5(pasaword string) string {
	hash := md5.New()
	hash.Write([]byte(pasaword))
	passwordHash := hash.Sum(nil)
	// 将密码转换为16进制储存
	passwordHash16 := hex.EncodeToString(passwordHash)
	return passwordHash16
}

// 加盐值加密
func Md5Password(password, salt string) string {
	return Md5(password + salt)
}

// 解密
func ValidMd5Password(password, salt, dataPwd string) bool {
	return Md5Password(password, salt) == dataPwd
}

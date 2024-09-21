package protocol

import (
	"unicode"
)

// 用户名合法: 5-16位的字母和数组成
func IsUsernameValid(username string) bool {
	runes := []rune(username)
	cnt := len(runes)
	// 长度校验
	if !(cnt >= 5 && cnt <= 16) {
		return false
	}
	for _, char := range runes {
		if (!unicode.IsDigit(char)) && (!unicode.IsLetter(char)) {
			return false
		}
	}
	return true
}

// 密码合法: 6-24位的字母和数组成
func IsPasswordValid(password string) bool {
	runes := []rune(password)
	cnt := len(runes)
	// 长度校验
	if !(cnt >= 6 && cnt <= 24) {
		return false
	}
	for _, char := range runes {
		if (!unicode.IsDigit(char)) && (!unicode.IsLetter(char)) {
			return false
		}
	}
	return true
}

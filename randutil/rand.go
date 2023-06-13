package randutil

import (
	"math/rand"
)

// String 随机字符，字母大小写+数字
func String(length int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	result := make([]rune, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// Number 随机字符，只包含数字
func Number(length int) string {
	result := make([]rune, length)
	for i := range result {
		result[i] = rune(rand.Intn(10) + 48)
	}
	return string(result)
}

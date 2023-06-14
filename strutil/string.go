package strutil

import (
	"math/rand"
	"strings"
	"unicode/utf8"
)

// Len 字符串长度, utf8.RuneCountInString() 别名
func Len(s string) int { return utf8.RuneCountInString(s) }

// Split 切割字符串，去除前后空格
func Split(s, sep string) []string {
	arr := strings.Split(s, sep)
	ret := make([]string, 0, len(arr))
	for _, v := range arr {
		if v = strings.TrimSpace(v); v != "" {
			ret = append(ret, v)
		}
	}
	return ret
}

// Random 生成数字+小写字母的随机字符串
func Random(n int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	result := make([]rune, n)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// RandomNumber 生成只有数字的随机字符串
func RandomNumber(n int) string {
	result := make([]rune, n)
	for i := range result {
		result[i] = rune(rand.Intn(10) + 48)
	}
	return string(result)
}

package fsutil

import "os"

// IsFile 检查是否文件
func IsFile(filename string) bool {
	info, err := os.Stat(filename)
	if err == nil && !info.IsDir() {
		return true
	}
	return false
}

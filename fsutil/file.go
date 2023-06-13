package fsutil

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// IsFile 检查是否文件
func IsFile(filename string) bool {
	info, err := os.Stat(filename)
	if err == nil && !info.IsDir() {
		return true
	}
	return false
}

// MD5File md5文件
func MD5File(fn string) (string, error) {
	file, err := os.Open(fn)
	if err != nil {
		return "", err
	}
	defer file.Close()
	const bufferSize = 10 << 20 // 10M
	buffer := make([]byte, bufferSize)
	reader := bufio.NewReader(file)
	hash := md5.New()
	for {
		n, err := reader.Read(buffer)
		if err == nil || err == io.EOF {
			_, _ = hash.Write(buffer[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

package hashutil

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
)

// MD5 计算MD5
func MD5(b []byte) string {
	hash := md5.New()
	_, err := hash.Write(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// MD5String 计算MD5
func MD5String(s string) string {
	return MD5([]byte(s))
}

// CRC32 计算CRC32
func CRC32(b []byte) uint32 {
	crc := crc32.NewIEEE()
	_, err := crc.Write(b)
	if err != nil {
		return 0
	}

	return crc.Sum32()
}

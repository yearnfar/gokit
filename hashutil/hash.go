package hashutil

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
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

// SHA1 计算SHA1哈希
func SHA1(b []byte) string {
	h := sha1.New()
	_, err := h.Write(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA256 计算SHA256哈希
func SHA256(b []byte) string {
	h := sha256.New()
	_, err := h.Write(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA512 计算SHA512哈希
func SHA512(b []byte) string {
	h := sha512.New()
	_, err := h.Write(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

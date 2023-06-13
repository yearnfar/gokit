package mathutil

// Number 包含所有数字类型
type Number interface {
	~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64 | ~float32 | ~float64
}

// Min 返回最小值
func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max 返回最大值
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

package mathutil

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
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

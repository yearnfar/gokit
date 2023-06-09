package arrutil

// InArray 判断数据是否在数组中
func InArray[T comparable](v T, arr []T) bool {
	for _, a := range arr {
		if a == v {
			return true
		}
	}
	return false
}

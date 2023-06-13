package arrutil

import (
	"fmt"
)

// InArray 判断数据是否在数组中
func InArray[T comparable](v T, arr []T) bool {
	for _, a := range arr {
		if a == v {
			return true
		}
	}
	return false
}

// NumbersToStrs 任意类型slice转为字符串数组
func NumbersToStrs[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](arr []T) []string {
	ret := make([]string, len(arr))
	for i := range arr {
		ret[i] = fmt.Sprint(arr[i])
	}
	return ret
}

// Merge 合并两个数组
func Merge[T any](arr1 []T, arr2 []T) []T {
	newArr := make([]T, len(arr1)+len(arr2))
	copy(newArr, arr1)
	copy(newArr[len(arr1):], arr2)
	return newArr
}

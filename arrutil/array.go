package arrutil

import (
	"fmt"

	"github.com/yearnfar/gokit/mathutil"
	"golang.org/x/exp/constraints"
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
func NumbersToStrs[T constraints.Integer | constraints.Float](arr []T) []string {
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

// Unique 去重
func Unique[T comparable](arr []T) []T {
	ret := make([]T, 0, len(arr))
	retMap := make(map[T]struct{})
	for _, a := range arr {
		if _, ok := retMap[a]; !ok {
			ret = append(ret, a)
			retMap[a] = struct{}{}
		}
	}
	return ret
}

// Intersect 两个数组取交集
func Intersect[T comparable](arr1 []T, arr2 []T) []T {
	ret := make([]T, 0, mathutil.Min(len(arr1), len(arr2)))
	for _, a := range arr1 {
		if InArray(a, arr2) {
			ret = append(ret, a)
		}
	}
	return ret
}

// Reverse 反转数组
func Reverse[T any](arr []T) []T {
	if len(arr) >= 2 {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}

// Split 将数组分为n等份
func Split[T any](arr []T, n int) (ret [][]T) {
	if n == 0 {
		return
	}
	size := len(arr) / n
	left := len(arr) % n

	ret = make([][]T, n)
	offset := 0

	for i := 0; i < n; i++ {
		delta := size
		if i < left {
			delta++
		}
		ret[i] = arr[offset : offset+delta]
		offset += delta
	}
	return ret
}

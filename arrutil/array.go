package arrutil

import (
	"fmt"

	"github.com/yearnfar/gokit/mathutil"
	"golang.org/x/exp/constraints"
)

// Max 数组取最大值
func Max[T constraints.Ordered](arr []T) (ret T) {
	if len(arr) == 0 {
		return
	}
	ret = arr[0]
	for _, v := range arr[1:] {
		if ret < v {
			ret = v
		}
	}
	return ret
}

// Min 数组取最小值
func Min[T constraints.Ordered](arr []T) (ret T) {
	if len(arr) == 0 {
		return
	}
	ret = arr[0]
	for _, v := range arr[1:] {
		if ret > v {
			ret = v
		}
	}
	return ret
}

// Column 返回数组对象中的int类型的列组成的数组
func Column[T, V any](arr []T, fn func(item T) V) (result []V) {
	for _, v := range arr {
		result = append(result, fn(v))
	}
	return
}

// ColumnMap 返回数组对象属性值映射的map
func ColumnMap[K comparable, V, T any](arr []T, fn func(item T) (K, V)) (result map[K]V) {
	result = make(map[K]V)
	for _, v := range arr {
		k, v := fn(v)
		result[k] = v
	}
	return
}

// ColumnUnique 返回数组对象中的int类型的列组成的数组并去重
func ColumnUnique[T any, V comparable](arr []T, fn func(item T) V) (result []V) {
	m := make(map[V]struct{})
	for _, v := range arr {
		val := fn(v)
		if _, ok := m[val]; !ok {
			result = append(result, val)
			m[val] = struct{}{}
		}
	}
	return
}

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
func Intersect[T comparable](arr1, arr2 []T) []T {
	ret := make([]T, 0, mathutil.Min(len(arr1), len(arr2)))
	for _, a := range arr1 {
		if InArray(a, arr2) {
			ret = append(ret, a)
		}
	}
	return ret
}

// Merge 合并两个数组
func Merge[T any](arr1, arr2 []T) []T {
	arr := make([]T, len(arr1)+len(arr2))
	copy(arr, arr1)
	copy(arr[len(arr1):], arr2)
	return arr
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

// Split 将数组分为n份
func Split[T any](arr []T, n int) (ret [][]T) {
	if n < 2 {
		return [][]T{arr}
	}
	size, left := len(arr)/n, len(arr)%n
	ret = make([][]T, n)
	offset := 0
	for i := 0; i < n; i++ {
		delta := size
		if i < left {
			delta++
		}
		ret[i] = make([]T, delta)
		copy(ret[i], arr[offset:offset+delta])
		offset += delta
	}
	return ret
}

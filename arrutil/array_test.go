package arrutil

import "testing"

func TestInArray(t *testing.T) {
	type testData[T comparable] struct {
		arr []T
		v   T
		b   bool
	}

	test1 := []testData[int]{
		{[]int{1, 2, 3}, 2, true},
		{[]int{111, 2222, 3333}, 3333, true},
		{[]int{1, 2, 3}, 6, false},
		{[]int{123, 122, 124}, 23, false},
	}
	for _, item := range test1 {
		if InArray(item.v, item.arr) != item.b {
			t.Fatalf("arr: %v, v: %v, export: %v", item.arr, item.v, item.b)
		}
	}

	test2 := []testData[string]{
		{[]string{"a", "b", "c"}, "a", true},
		{[]string{"你好", "哈哈", "哦"}, "哈哈", true},
		{[]string{"1", "2", "3"}, "6", false},
		{[]string{"hello world", "thank you", "xasfd"}, "嘿", false},
	}
	for _, item := range test2 {
		if InArray(item.v, item.arr) != item.b {
			t.Fatalf("arr: %v, v: %v, export: %v", item.arr, item.v, item.b)
		}
	}
	t.Log("done")
}

package arrutil

import (
	"strings"
	"testing"
)

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

func TestNumbersToString(t *testing.T) {
	testData1 := []struct {
		arr    []int
		expect []string
	}{
		{[]int{1, 2, 3, 4, 5}, []string{"1", "2", "3", "4", "5"}},
		{[]int{11, 22, 33, 44, 55}, []string{"11", "22", "33", "44", "55"}},
	}
	for _, item := range testData1 {
		out := NumbersToStrs(item.arr)
		if strings.Join(out, ",") != strings.Join(item.expect, ",") {
			t.Fatalf("number : %v , export: %v", item.arr, item.expect)
		}
	}
}

func TestMerge(t *testing.T) {
	testData1 := [][3][]string{
		{{"你"}, {"好"}, {"你", "好"}},
		{{"再"}, {"见"}, {"再", "见"}},
		{{"安和"}, {"桥"}, {"安和", "桥"}},
	}

	for _, item := range testData1 {
		out := Merge(item[0], item[1])
		if strings.Join(out, ",") != strings.Join(item[2], ",") {
			t.Fatalf("merge : %v and %v, export: %v", item[0], item[1], item[2])
		}
	}

	testData2 := [][3][]int{
		{{1}, {3, 4}, {1, 3, 4}},
		{{2, 3}, {5}, {2, 3, 5}},
		{{1, 2}, {3, 4}, {1, 2, 3, 4}},
	}

	for _, item := range testData2 {
		out := Merge(item[0], item[1])
		if strings.Join(NumbersToStrs(out), ",") != strings.Join(NumbersToStrs(item[2]), ",") {
			t.Fatalf("merge : %v and %v, export: %v", item[0], item[1], item[2])
		}
	}
}

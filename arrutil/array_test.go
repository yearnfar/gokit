package arrutil

import (
	"fmt"
	"testing"
)

func TestInArray(t *testing.T) {
	testData := []struct {
		arr    []interface{}
		v      interface{}
		expect bool
	}{
		{[]interface{}{1, 2, 3}, 2, true},
		{[]interface{}{111, 2222, 3333}, 3333, true},
		{[]interface{}{1, 2, 3}, 6, false},
		{[]interface{}{123, 122, 124}, 23, false},
		{[]interface{}{"a", "b", "c"}, "a", true},
		{[]interface{}{"你好", "哈哈", "哦"}, "哈哈", true},
		{[]interface{}{"1", "2", "3"}, "6", false},
		{[]interface{}{"hello world", "thank you", "xasfd"}, "嘿", false},
	}

	for _, item := range testData {
		ret := InArray(item.v, item.arr)
		if ret != item.expect {
			t.Fatalf("arr: %v, v: %v, got: %v, export: %v", item.arr, item.v, ret, item.expect)
		}
	}
}

func TestNumbersToString(t *testing.T) {
	testData := []struct {
		arr    interface{}
		expect []string
	}{
		{[]int{1, 2, 3, 4, 5}, []string{"1", "2", "3", "4", "5"}},
		{[]int{11, 22, 33, 44, 55}, []string{"11", "22", "33", "44", "55"}},
		{[]int{1, 2, 3, 4, 5}, []string{"1", "2", "3", "4", "5"}},
		{[]float64{1.1, 2.1, 3.0, 4.0, 5.0}, []string{"1.1", "2.1", "3.0", "4.0", "5.0"}},
	}

	for _, item := range testData {
		switch value := item.arr.(type) {
		case []int:
			ret := NumbersToStrs(value)
			if fmt.Sprint(ret) != fmt.Sprint(item.expect) {
				t.Fatalf("number : %v , export: %v", item.arr, item.expect)
			}
		case []float64:
			ret := NumbersToStrs(value)
			if fmt.Sprint(ret) != fmt.Sprint(item.expect) {
				t.Fatalf("number : %v , export: %v", item.arr, item.expect)
			}
		case []float32:
			ret := NumbersToStrs(value)
			if fmt.Sprint(ret) != fmt.Sprint(item.expect) {
				t.Fatalf("number : %v , export: %v", item.arr, item.expect)
			}
		}
	}
}

func TestMerge(t *testing.T) {
	testData := []struct {
		arr    [2][]interface{}
		expect []interface{}
	}{
		{[2][]interface{}{{"你"}, {"好"}}, []interface{}{"你", "好"}},
		{[2][]interface{}{{"再"}, {"见"}}, []interface{}{"再", "见"}},
		{[2][]interface{}{{"安和"}, {"桥"}}, []interface{}{"安和", "桥"}},
		{[2][]interface{}{{1}, {3, 4}}, []interface{}{1, 3, 4}},
		{[2][]interface{}{{2, 3}, {5}}, []interface{}{2, 3, 5}},
		{[2][]interface{}{{1, 2}, {3, 4}}, []interface{}{1, 2, 3, 4}},
	}

	for _, item := range testData {
		ret := Merge(item.arr[0], item.arr[1])
		if fmt.Sprint(ret) != fmt.Sprint(item.expect) {
			t.Fatalf("merge : %v and %v, ret: %v, export: %v", item.arr[0], item.arr[1], ret, item.expect)
		}
	}
}

func TestSplit(t *testing.T) {
	testData1 := []struct {
		arr    []interface{}
		n      int
		expect [][]interface{}
	}{
		{[]interface{}{1, 2, 3, 4, 5}, 0, [][]interface{}{}},
		{[]interface{}{1, 2, 3, 4, 5}, 1, [][]interface{}{{1, 2, 3, 4, 5}}},
		{[]interface{}{1, 2, 3, 4, 5}, 2, [][]interface{}{{1, 2, 3}, {4, 5}}},
		{[]interface{}{1, 2, 3, 4, 5}, 3, [][]interface{}{{1, 2}, {3, 4}, {5}}},
		{[]interface{}{1, 2, 3, 4, 5}, 4, [][]interface{}{{1, 2}, {3}, {4}, {5}}},
		{[]interface{}{1, 2, 3, 4, 5}, 5, [][]interface{}{{1}, {2}, {3}, {4}, {5}}},
		{[]interface{}{1, 2, 3, 4, 5}, 6, [][]interface{}{{1}, {2}, {3}, {4}, {5}, {}}},
		{[]interface{}{"你", "好", "再", "见"}, 0, [][]interface{}{}},
		{[]interface{}{"你", "好", "再", "见"}, 1, [][]interface{}{{"你", "好", "再", "见"}}},
		{[]interface{}{"你", "好", "再", "见"}, 2, [][]interface{}{{"你", "好"}, {"再", "见"}}},
		{[]interface{}{"你", "好", "再", "见"}, 3, [][]interface{}{{"你", "好"}, {"再"}, {"见"}}},
		{[]interface{}{"你", "好", "再", "见"}, 4, [][]interface{}{{"你"}, {"好"}, {"再"}, {"见"}}},
		{[]interface{}{"你", "好", "再", "见"}, 5, [][]interface{}{{"你"}, {"好"}, {"再"}, {"见"}, {}}},
	}

	for _, item := range testData1 {
		ret := Split(item.arr, item.n)
		if fmt.Sprint(ret) != fmt.Sprint(item.expect) {
			t.Fatalf("arr: %v, n: %d, got: %v. expect: %v", item.arr, item.n, ret, item.expect)
		}
	}
}

package arrutil

import (
	"reflect"
	"testing"
)

func TestArrayColumn(t *testing.T) {
	type Object struct {
		Id   any
		Name string
	}
	testData := map[string][]Object{
		"ints":   {{1, "a"}, {2, "b"}, {3, "c"}, {4, "d"}, {4, ""}},
		"strs":   {{"1", "a"}, {"2", "b"}, {"3", "c"}, {"4", "d"}, {"4", ""}},
		"floats": {{1.0, "a"}, {2.0, "b"}, {3.0, "c"}, {4.0, "d"}, {4.0, ""}},
	}
	for typ, list := range testData {
		switch typ {
		case "ints":
			ids := Column(list, func(item Object) int { return item.Id.(int) })
			want := []int{1, 2, 3, 4, 4}
			if !reflect.DeepEqual(ids, want) {
				t.Errorf("Column() = %v, want %v", ids, want)
			}
		case "strs":
			ids := Column(list, func(item Object) string { return item.Id.(string) })
			want := []string{"1", "2", "3", "4", "4"}
			if !reflect.DeepEqual(ids, want) {
				t.Errorf("Column() = %v, want %v", ids, want)
			}
		case "floats":
			ids := Column(list, func(item Object) float64 { return item.Id.(float64) })
			want := []float64{1.0, 2.0, 3.0, 4.0, 4.0}
			if !reflect.DeepEqual(ids, want) {
				t.Errorf("Column() = %v, want %v", ids, want)
			}
		}
	}
}

func TestArrayColumnMap(t *testing.T) {
	type Object struct {
		Id   any
		Name string
	}
	testData := map[string][]Object{
		"ints":   {{1, "a"}, {2, "b"}, {3, "c"}, {4, "d"}, {4, ""}},
		"strs":   {{"1", "a"}, {"2", "b"}, {"3", "c"}, {"4", "d"}, {"4", ""}},
		"floats": {{1.0, "a"}, {2.0, "b"}, {3.0, "c"}, {4.0, "d"}, {4.0, ""}},
	}
	for typ, list := range testData {
		switch typ {
		case "ints":
			ids := ColumnMap(list, func(item Object) (int, string) { return item.Id.(int), item.Name })
			want := map[int]string{1: "a", 2: "b", 3: "c", 4: ""}
			if !reflect.DeepEqual(ids, want) {
				t.Errorf("ColumnMap() = %v, want %v", ids, want)
			}
		case "strs":
			ids := ColumnMap(list, func(item Object) (string, string) { return item.Id.(string), item.Name })
			want := map[string]string{"1": "a", "2": "b", "3": "c", "4": ""}
			if !reflect.DeepEqual(ids, want) {
				t.Errorf("ColumnMap() = %v, want %v", ids, want)
			}
		case "floats":
			out := ColumnMap(list, func(item Object) (float64, string) { return item.Id.(float64), item.Name })
			want := map[float64]string{1.0: "a", 2.0: "b", 3.0: "c", 4.0: ""}
			if !reflect.DeepEqual(out, want) {
				t.Errorf("ColumnMap() = %v, want %v", out, want)
			}
		}
	}
}

func TestArrayColumnUniq(t *testing.T) {
	type Object struct {
		Id   any
		Name string
	}
	testData := map[string][]Object{
		"ints":   {{1, "a"}, {1, "b"}, {2, "c"}, {2, "d"}, {3, ""}},
		"strs":   {{"1", "a"}, {"1", "b"}, {"2", "c"}, {"2", "d"}, {"3", ""}},
		"floats": {{1.0, "a"}, {1.0, "b"}, {2.0, "c"}, {2.0, "d"}, {3.0, ""}},
	}
	for typ, list := range testData {
		switch typ {
		case "ints":
			ids := ColumnUnique(list, func(item Object) int { return item.Id.(int) })
			want := []int{1, 2, 3}
			if !reflect.DeepEqual(ids, want) {
				t.Errorf("Column() = %v, want %v", ids, want)
			}
		case "strs":
			ids := ColumnUnique(list, func(item Object) string { return item.Id.(string) })
			want := []string{"1", "2", "3"}
			if !reflect.DeepEqual(ids, want) {
				t.Errorf("Column() = %v, want %v", ids, want)
			}
		case "floats":
			ids := ColumnUnique(list, func(item Object) float64 { return item.Id.(float64) })
			want := []float64{1.0, 2.0, 3.0}
			if !reflect.DeepEqual(ids, want) {
				t.Errorf("Column() = %v, want %v", ids, want)
			}
		}
	}
}

func TestInArray(t *testing.T) {
	testData := []struct {
		arr    []any
		v      any
		expect bool
	}{
		{[]any{1, 2, 3}, 2, true},
		{[]any{111, 2222, 3333}, 3333, true},
		{[]any{1, 2, 3}, 6, false},
		{[]any{123, 122, 124}, 23, false},
		{[]any{"a", "b", "c"}, "a", true},
		{[]any{"你好", "哈哈", "哦"}, "哈哈", true},
		{[]any{"1", "2", "3"}, "6", false},
		{[]any{"hello world", "thank you", "xasfd"}, "嘿", false},
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
		arr    any
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
			if ret := NumbersToStrs(value); !reflect.DeepEqual(ret, item.expect) {
				t.Fatalf("number : %v , export: %v", item.arr, item.expect)
			}
		case []float64:
			if ret := NumbersToStrs(value); !reflect.DeepEqual(ret, item.expect) {
				t.Fatalf("number : %v , export: %v", item.arr, item.expect)
			}
		case []float32:
			if ret := NumbersToStrs(value); !reflect.DeepEqual(ret, item.expect) {
				t.Fatalf("number : %v , export: %v", item.arr, item.expect)
			}
		}
	}
}

func TestMerge(t *testing.T) {
	testData := []struct {
		arr    [2][]any
		expect []any
	}{
		{[2][]any{{"你"}, {"好"}}, []any{"你", "好"}},
		{[2][]any{{"再"}, {"见"}}, []any{"再", "见"}},
		{[2][]any{{"安和"}, {"桥"}}, []any{"安和", "桥"}},
		{[2][]any{{1}, {3, 4}}, []any{1, 3, 4}},
		{[2][]any{{2, 3}, {5}}, []any{2, 3, 5}},
		{[2][]any{{1, 2}, {3, 4}}, []any{1, 2, 3, 4}},
	}

	for _, item := range testData {
		ret := Merge(item.arr[0], item.arr[1])
		if !reflect.DeepEqual(ret, item.expect) {
			t.Fatalf("merge : %v and %v, ret: %v, export: %v", item.arr[0], item.arr[1], ret, item.expect)
		}
	}
}

func TestSplit(t *testing.T) {
	testData1 := []struct {
		arr    []any
		n      int
		expect [][]any
	}{
		{[]any{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1, [][]any{{1, 2, 3, 4, 5, 6, 7, 8, 9}}},
		{[]any{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2, [][]any{{1, 2, 3, 4, 5}, {6, 7, 8, 9}}},
		{[]any{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, [][]any{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		{[]any{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4, [][]any{{1, 2, 3}, {4, 5}, {6, 7}, {8, 9}}},
		{[]any{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, [][]any{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}}},
		{[]any{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6, [][]any{{1, 2}, {3, 4}, {5, 6}, {7}, {8}, {9}}},
		{[]any{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7, [][]any{{1, 2}, {3, 4}, {5}, {6}, {7}, {8}, {9}}},
		{[]any{1, 2, 3, 4, 5, 6, 7, 8, 9}, 8, [][]any{{1, 2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}},
		{[]any{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, [][]any{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}},
		{[]any{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, [][]any{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {}}},
		{[]any{"你", "好", "再", "见"}, 0, [][]any{{"你", "好", "再", "见"}}},
		{[]any{"你", "好", "再", "见"}, 1, [][]any{{"你", "好", "再", "见"}}},
		{[]any{"你", "好", "再", "见"}, 2, [][]any{{"你", "好"}, {"再", "见"}}},
		{[]any{"你", "好", "再", "见"}, 3, [][]any{{"你", "好"}, {"再"}, {"见"}}},
		{[]any{"你", "好", "再", "见"}, 4, [][]any{{"你"}, {"好"}, {"再"}, {"见"}}},
		{[]any{"你", "好", "再", "见"}, 5, [][]any{{"你"}, {"好"}, {"再"}, {"见"}, {}}},
	}

	for _, item := range testData1 {
		ret := Split(item.arr, item.n)
		if !reflect.DeepEqual(ret, item.expect) {
			t.Fatalf("arr: %v, n: %d, got: %v. expect: %v", item.arr, item.n, ret, item.expect)
		}
	}
}

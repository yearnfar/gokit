package fsutil

import (
	"testing"
)

func TestHumanReadableSize(t *testing.T) {
	var testData = []struct {
		size   int64
		expect string
	}{
		{1 << 10, "1.00 KB"},
		{1 << 20, "1.00 MB"},
		{1 << 30, "1.00 GB"},
		{1 << 40, "1.00 TB"},
		{1 << 50, "1.00 PB"},
		{1 << 60, "1.00 EB"},
	}

	for _, item := range testData {
		ret := HumanReadableSize(item.size)
		if ret != item.expect {
			t.Fatalf("size: %d, get: %s, expect: %s", item.size, ret, item.expect)
		}
	}
}

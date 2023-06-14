package datetime

import (
	"testing"
	"time"
)

func TestStrftime(t *testing.T) {
	testData := []struct {
		t      time.Time
		layout string
		expect string
	}{
		{time.Date(2023, 6, 14, 0, 0, 0, 0, time.Local), "Y-m-d H:i:s", "2023-06-14 00:00:00"},
		{time.Date(2023, 6, 6, 23, 0, 0, 0, time.Local), "y-n-j h:i:s", "23-6-6 11:00:00"},
	}

	for _, item := range testData {
		ret := Strftime(item.t, item.layout)
		if ret != item.expect {
			t.Fatalf("t: %v, get: %s, expect: %s", item.t, ret, item.expect)
		}
	}
}

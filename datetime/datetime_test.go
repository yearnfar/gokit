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
		// 年
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%Y", "2006"},
		{time.Date(2023, 01, 02, 15, 04, 05, 0, time.Local), "%Y", "2023"},
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%-Y", "06"},
		{time.Date(2023, 01, 02, 15, 04, 05, 0, time.Local), "%-Y", "23"},
		// 月
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%m", "01"},
		{time.Date(2006, 12, 02, 15, 04, 05, 0, time.Local), "%m", "12"},
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%-m", "1"},
		{time.Date(2006, 12, 02, 15, 04, 05, 0, time.Local), "%-m", "12"},
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%b", "Jan"},
		{time.Date(2006, 12, 02, 15, 04, 05, 0, time.Local), "%b", "Dec"},
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%B", "January"},
		{time.Date(2006, 12, 02, 15, 04, 05, 0, time.Local), "%B", "December"},
		// 日
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%d", "02"},
		{time.Date(2006, 01, 12, 15, 04, 05, 0, time.Local), "%d", "12"},
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%-d", "2"},
		{time.Date(2006, 01, 12, 15, 04, 05, 0, time.Local), "%-d", "12"},
		// 时
		{time.Date(2006, 01, 02, 01, 04, 05, 0, time.Local), "%H", "01"},
		{time.Date(2006, 01, 02, 12, 04, 05, 0, time.Local), "%H", "12"},
		{time.Date(2006, 01, 02, 23, 04, 05, 0, time.Local), "%H", "23"},
		{time.Date(2006, 01, 02, 01, 04, 05, 0, time.Local), "%I", "01"},
		{time.Date(2006, 01, 02, 12, 04, 05, 0, time.Local), "%I", "12"},
		{time.Date(2006, 01, 02, 23, 04, 05, 0, time.Local), "%I", "11"},
		{time.Date(2006, 01, 02, 01, 04, 05, 0, time.Local), "%-I", "1"},
		{time.Date(2006, 01, 02, 12, 04, 05, 0, time.Local), "%-I", "12"},
		{time.Date(2006, 01, 02, 23, 04, 05, 0, time.Local), "%-I", "11"},
		// 分
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%M", "04"},
		{time.Date(2006, 01, 02, 15, 34, 05, 0, time.Local), "%M", "34"},
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%-M", "4"},
		{time.Date(2006, 01, 02, 15, 34, 05, 0, time.Local), "%-M", "34"},
		// 秒
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%S", "05"},
		{time.Date(2006, 01, 02, 15, 04, 35, 0, time.Local), "%S", "35"},
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%-S", "5"},
		{time.Date(2006, 01, 02, 15, 04, 35, 0, time.Local), "%-S", "35"},
		// 周
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%a", "Mon"},
		{time.Date(2008, 01, 02, 15, 04, 05, 0, time.Local), "%a", "Wed"},
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%A", "Monday"},
		{time.Date(2008, 01, 02, 15, 04, 05, 0, time.Local), "%A", "Wednesday"},
		// AM、PM
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%p", "PM"},
		{time.Date(2006, 01, 02, 03, 04, 05, 0, time.Local), "%p", "AM"},
		// 地区
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%Z", "CST"},
		{time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local), "%z", "+0800"},
	}

	for _, item := range testData {
		ret := Strftime(item.t, item.layout)
		if ret != item.expect {
			t.Fatalf("t: %v, get: %s, expect: %s", item.t, ret, item.expect)
		}
	}
}

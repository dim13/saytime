package saytime

import (
	"testing"
	"time"
)

var testCases = []struct {
	time string
	text string
}{
	{"01:05AM", "The time is now, exactly five past one, in the morning"},
	{"02:11AM", "The time is now, just after ten past two, in the morning"},
	{"03:17AM", "The time is now, a little after quarter past three, in the morning"},
	{"04:18AM", "The time is now, almost twenty past four, in the morning"},

	{"05:25AM", "The time is now, exactly twenty-five past five, in the morning"},
	{"06:31AM", "The time is now, just after half past six, in the morning"},
	{"06:37AM", "The time is now, a little after twenty-five to seven, in the morning"},
	{"07:38AM", "The time is now, almost twenty to eight, in the morning"},

	{"08:45AM", "The time is now, exactly quarter to nine, in the morning"},
	{"09:51AM", "The time is now, just after ten to ten, in the morning"},
	{"10:57AM", "The time is now, a little after five to eleven, in the morning"},
	{"11:58AM", "The time is now, almost twelve"},

	{"12:56PM", "The time is now, just after five to one, in the afternoon"},
	{"01:52PM", "The time is now, a little after ten to two, in the afternoon"},
	{"02:45PM", "The time is now, exactly quarter to three, in the afternoon"},
	{"03:39PM", "The time is now, almost twenty to four, in the afternoon"},

	{"04:36PM", "The time is now, just after twenty-five to five, in the afternoon"},
	{"06:32PM", "The time is now, a little after half past six, in the evening"},
	{"07:25PM", "The time is now, exactly twenty-five past seven, in the evening"},
	{"08:19PM", "The time is now, almost twenty past eight, in the evening"},

	{"09:16PM", "The time is now, just after quarter past nine, in the evening"},
	{"10:09PM", "The time is now, almost ten past ten, in the evening"},
	{"11:05PM", "The time is now, exactly five past eleven, in the evening"},
	{"11:47PM", "The time is now, a little after quarter to midnight"},
}

func TestSayTime(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.time, func(t *testing.T) {
			clock, err := time.Parse(time.Kitchen, tc.time)
			if err != nil {
				t.Fatal(err)
			}
			text := Time{clock}.String()
			if text != tc.text {
				t.Errorf("got %q, want %q", text, tc.text)
			}
		})
	}
}

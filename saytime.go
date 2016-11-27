// Package saytime translates time into plain English
package saytime

import "time"

// Time wraps time.Time
type Time struct {
	time.Time
}

var approx = []string{
	"exactly", "just after", "a little after", "almost", "almost",
}

func (n Time) approx() string {
	return approx[n.Minute()%len(approx)]
}

var minutes = []string{
	"", " five past", " ten past", " quarter past",
	" twenty past", " twenty-five past", " half past", " twenty-five to",
	" twenty to", " quarter to", " ten to", " five to",
}

func (n Time) minute() string {
	m := n.Minute() / 5
	if n.Minute()%5 > 2 {
		m++
	}
	return minutes[m%len(minutes)]
}

var hours = []string{
	" twelve", " one", " two", " three", " four", " five",
	" six", " seven", " eight", " nine", " ten", " eleven",
}

func (n Time) hour() string {
	h, m := n.Hour(), n.Minute()
	if m > 33 {
		h++
	}
	if h == 24 {
		return " midnight"
	}
	return hours[h%len(hours)]
}

func (n Time) tod() string {
	h, m := n.Hour(), n.Minute()
	switch {
	case h == 11 && m > 58, h == 23 && m > 33:
		return "."
	case h > 17:
		return ", in the evening"
	case h > 11:
		return ", in the afternoon"
	default:
		return ", in the morning"
	}
}

func (n Time) String() string {
	return "The time is now, " + n.approx() + n.minute() + n.hour() + n.tod()
}

// Now represents time now
func Now() Time {
	return Time{time.Now()}
}

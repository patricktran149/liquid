package values

import (
	"reflect"
	"time"
)

var zeroTime time.Time

var dateLayouts = []string{
	// from the Go library
	time.ANSIC,    // "Mon Jan _2 15:04:05 2006"
	time.UnixDate, // "Mon Jan _2 15:04:05 MST 2006"
	time.RubyDate, // "Mon Jan 02 15:04:05 -0700 2006"
	time.RFC822,   // "02 Jan 06 15:04 MST"
	time.RFC822Z,  // "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	time.RFC850,   // "Monday, 02-Jan-06 15:04:05 MST"
	time.RFC1123,  // "Mon, 02 Jan 2006 15:04:05 MST"
	time.RFC1123Z, // "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	time.RFC3339,  // "2006-01-02T15:04:05Z07:00"

	// ISO 8601
	"2006-01-02T15:04:05-07:00", // this is also XML Schema
	"2006-01-02T15:04:05Z",
	"2006-01-02",
	"20060102T150405Z",

	// from Ruby's Time.parse docs
	"Mon, 02 Jan 2006 15:04:05 -0700", // "RFC822" -- but not really

	// From Jekyll docs
	"02 January 2006", // Jekyll long string
	"02 Jan 2006",     // Jekyll short string

	// observed in the wild; plus some variants
	"2006-01-02 15:04:05 -07:00",
	"2006-01-02 15:04:05 -0700",
	"2006-01-02 15:04:05 MST",
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"January 2, 2006",
	"January 2 2006",
	"Jan 2, 2006",
	"Jan 2 2006",
}

// ParseDate tries a few heuristics to parse a date from a string
func ParseDate(s string) (time.Time, error) {
	if s == "now" {
		return time.Now(), nil
	}
	for _, layout := range dateLayouts {
		t, err := time.ParseInLocation(layout, s, time.Local)
		if err == nil {
			return t, nil
		}
	}
	return zeroTime, conversionError("", s, reflect.TypeOf(zeroTime))
}

// PHO ADD FUNCTION HERE
func ParseDateFromUnix(i int) (time.Time, error) {
	return time.Unix(int64(i), 0), nil
}

// PHO ADD FUNCTION HERE
func ParseDateFromUnixInt64(i int64) (time.Time, error) {
	return time.Unix(i, 0), nil
}

// PHO ADD FUNCTION HERE
func ParseDateFromUnixFloat64(f float64) (time.Time, error) {
	return time.Unix(int64(f), 0), nil
}

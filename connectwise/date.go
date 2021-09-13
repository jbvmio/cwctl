package connectwise

import "time"

const (
	DateLayout = `2006-01-02T15:04:05`
)

// Date works with the formatting of dates for the CW API.
type Date string

// MustGet parses and returns the time.Time for the CW formatted date or panics.
func (d Date) MustGet() time.Time {
	t, err := time.Parse(DateLayout, string(d))
	if err != nil {
		panic(err)
	}
	return t
}

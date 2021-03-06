// Package date contains methods and structs to deal with timeless date
package date

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"
)

// ErrMsgInvalidFormat reprensents the error message returned when an invalid
// update format is provided
var ErrMsgInvalidFormat = "invalid format"

// DATE is a time.Time layout for the a date (no time)
const DATE = "2006-01-02"

// Date represents a time.Time that uses DATE for json input/output
// instead of RFC3339
type Date struct {
	time.Time
}

// Today returns the current local day.
func Today() *Date {
	var day, year int
	var month time.Month
	year, month, day = time.Now().UTC().Date()
	return &Date{Time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

// New accepts "year-month" or "year-month-day"
func New(date string) (Date, error) {
	// If we only have year-month, then we add "-day"
	if strings.Count(date, "-") == 1 {
		date += "-01"
	}

	t, err := time.Parse(DATE, date)
	if err != nil {
		return Date{}, err
	}
	return Date{Time: t}, nil
}

// Value returns a value that the database can handle
// https://golang.org/pkg/database/sql/driver/#Valuer
func (t *Date) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return t.Format(DATE), nil
}

// Scan assigns a value from a database driver
// https://golang.org/pkg/database/sql/#Scanner
func (t *Date) Scan(value interface{}) error {
	if value != nil {
		t.Time = value.(time.Time)
	}
	return nil
}

// String implements the fmt.Stringer interface
// https://golang.org/pkg/fmt/#Stringer
func (t Date) String() string {
	return t.Format(DATE)
}

// ScanString implements the go-params Scanner interface
func (t *Date) ScanString(date string) error {
	if strings.Count(date, "-") == 1 {
		date += "-01"
	}

	var err error
	t.Time, err = time.Parse(DATE, date)
	if err != nil {
		return errors.New(ErrMsgInvalidFormat)
	}
	t.Time.UTC()
	return nil
}

// MarshalJSON returns a valid json representation of the struct
// https://golang.org/pkg/encoding/json/#Marshaler
func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format(DATE) + `"`), nil
}

// UnmarshalJSON tries to parse a json data into a valid struct
// https://golang.org/pkg/encoding/json/#Unmarshaler
func (t *Date) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return nil
	}

	t.Time, err = time.Parse(`"`+DATE+`"`, string(data))
	return
}

// Equal checks if the given date is equal to the current one
func (t Date) Equal(u Date) bool {
	return (t.Time.Year() == u.Time.Year()) &&
		(t.Time.Month() == u.Time.Month()) &&
		(t.Time.Day() == u.Time.Day())
}

// IsBefore checks if the current date is before the given one
func (t Date) IsBefore(u Date) bool {
	if t.Time.Year() != u.Time.Year() {
		return t.Time.Year() < u.Time.Year()
	}
	if t.Time.Month() != u.Time.Month() {
		return t.Time.Month() < u.Time.Month()
	}
	if t.Time.Day() != u.Time.Day() {
		return t.Time.Day() < u.Time.Day()
	}
	return false
}

// IsAfter checks if the current date is after the given one
func (t Date) IsAfter(u Date) bool {
	if t.Time.Year() != u.Time.Year() {
		return t.Time.Year() > u.Time.Year()
	}
	if t.Time.Month() != u.Time.Month() {
		return t.Time.Month() > u.Time.Month()
	}
	if t.Time.Day() != u.Time.Day() {
		return t.Time.Day() > u.Time.Day()
	}
	return false
}

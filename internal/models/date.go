package models

import (
	"errors"
	"strings"
	"time"
)

var ErrInvalidDateFormat = errors.New("invalid date format, use YYYY-MM-DD")

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	if s == "" {
		return ErrInvalidDateFormat
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return ErrInvalidDateFormat
	}

	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	return []byte(`"` + t.Format("2006-01-02") + `"`), nil
}

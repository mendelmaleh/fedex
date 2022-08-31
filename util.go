package fedex

import (
	"time"
)

type DateISO struct {
	time.Time
}

func (d *DateISO) UnmarshalJSON(data []byte) error {
	// unquote
	s := string(data[1 : len(data)-1])

	if len(s) == 0 {
		return nil
	}

	t, err := time.Parse("2006-01-02T15:04:05-07:00"[:len(s)], s)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

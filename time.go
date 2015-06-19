package vcapi

import (
	"database/sql/driver"
	"time"
)

// Veracross Time Format
const VCTimeFormat = "2006-01-02"

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		d.Time = time.Time{}
		return nil
	}
	d.Time, err = time.Parse(`"`+VCTimeFormat+`"`, string(data))
	return
}

func (d Date) Value() (driver.Value, error) {
	return d.Time, nil
}

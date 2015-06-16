package vcapi

import "time"

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

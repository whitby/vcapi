package vcapi

import "fmt"

type Attendance struct {
	Name               string
	Grade              string
	StudentID          int    `json:"person_fk"`
	Status             string `json:"status"`
	Today              bool   `json:"today"`
	AttendanceDate     string `json:"attendance_date"`
	UpdateDate         string `json:"update_date"`
	EarlyDismissalTime string `json:"early_dismissal_time"`
	ReturnTime         string `json:"return_time"`
}

type DayAttendance []Attendance

// Fetch Attendance for today.
func (s *DayAttendance) Fetch(url string) (*DayAttendance, error) {
	url = fmt.Sprintf("%v/attendance.json", url)
	err := fetch(url, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Len is part of sort.Interface
func (d DayAttendance) Len() int {
	return len(d)
}

// Swap is part of sort.Interface.
func (d DayAttendance) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (d DayAttendance) Less(i, j int) bool {
	return d[i].Grade < d[j].Grade
}

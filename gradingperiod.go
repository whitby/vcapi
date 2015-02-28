package vcapi

import "fmt"

type GradingPeriod struct {
	EndDate     string `json:"end_date"`
	StartDate   string `json:"start_date"`
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type GradingPeriods []GradingPeriod

// Fetch Grading Periods
func (s *GradingPeriods) Fetch(url string) (*GradingPeriods, error) {
	url = fmt.Sprintf("%v/grading_periods.json", url)
	err := Fetch(url, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

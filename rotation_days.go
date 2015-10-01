package vcapi

type RotationDayService struct {
	client *Client
}

const (
	rotationDaysPath = "rotation_days"
)

type RotationDays struct {
	CalendarDayPK int    `json:"calendar_day_pk"`
	Description   string `json:"description"`
	Categories    string `json:"categories"`
	Rotation      int    `json:"rotation"`
	StartDate     Date   `json:"start_date"`
	EndDate       Date   `json:"end_date"`
}

func (s RotationDayService) Rotation(rotation int) (*[]RotationDays, error) {
	// todo not implemented.
	// all our days are rotation 2
	return nil, nil
}

func (s RotationDayService) list(opt *ListOptions, basePath string) ([]RotationDays, error) {
	// build url
	path := addOptions(basePath, format, opt)

	var rotations = []RotationDays{}
	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, nil
	}
	resp, err := s.client.Do(req, &rotations)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// handle pagination
	paginate(resp, opt)

	return rotations, nil
}

func (s RotationDayService) List(opt *ListOptions) ([]RotationDays, error) {
	rotations, err := s.list(opt, rotationDaysPath)
	if err != nil {
		return nil, err
	}

	return rotations, nil
}

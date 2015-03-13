package vcapi

type Parent struct {
	Person
	Occupation string `json:"occupation"`
}

type Parents []Parent

// Fetch a list of Parents
func (s *Parents) Fetch(url string) error {
	err := Fetch(url, &s)
	if err != nil {
		return err
	}
	return nil
}

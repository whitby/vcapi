package vcapi

type Parent struct {
	Person
	Occupation string `json:"occupation"`
}

type Parents []Parent

// Fetch a list of Parents
func (s *Parents) Fetch(url string) (*Parents, error) {
	err := fetch(url, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

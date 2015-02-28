package vcapi

type Students []Student

type Student struct {
	Person
	Grade string `json:"current_grade"`
}

// Fetch a student
func (s *Student) Fetch(url string) (*Student, error) {
	type aStudent struct {
		Student `json:"student"`
	}
	var a aStudent

	err := fetch(url, &a)
	if err != nil {
		return nil, err
	}
	*s = a.Student
	return s, nil
}

// Fetch A list of students
func (s *Students) Fetch(url string) (*Students, error) {
	err := fetch(url, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
